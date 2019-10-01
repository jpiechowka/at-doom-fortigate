package doom

import (
	"at-doom-fortigate/config"
	"at-doom-fortigate/networking"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"sync"
)

func CleanAndParseResponses(responsesChan <-chan *networking.MiniResponseObject) <-chan []byte {
	cleanAndParsedResponsesChan := make(chan []byte, config.CleanAndParsedResponsesChanBufferSize)

	go func() {
		defer close(cleanAndParsedResponsesChan)
		wg := sync.WaitGroup{}

		for response := range responsesChan {
			wg.Add(1)

			go func(r *networking.MiniResponseObject) {
				defer wg.Done()

				if r.RequestError == nil && r.ResponseBody != nil {
					outputJson, jsonMarshalError := json.Marshal(r)

					if jsonMarshalError != nil {
						log.Error().Err(jsonMarshalError).Msg("Error encoding struct to JSON")
						return
					}

					responseToWrite := append(outputJson, "\n"...)
					cleanAndParsedResponsesChan <- responseToWrite
				}
			}(response)
		}
		wg.Wait()
	}()

	return cleanAndParsedResponsesChan
}
