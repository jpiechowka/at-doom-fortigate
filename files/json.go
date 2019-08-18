package files

import (
	"at-doom-fortigate/config"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
	"sync"
	"sync/atomic"
)

func ReadRapid7JsonWithTargets(jsonFilePath string) <-chan string {
	targetsChan := make(chan string, config.ParsedTargetsChanBufferSize)

	go func() {
		defer close(targetsChan)
		wg := sync.WaitGroup{}

		var targetCtr uint64 = 0

		for jsonFileLine := range readFileLines(jsonFilePath) {
			wg.Add(1)
			go func(json string) {
				defer wg.Done()

				ip := gjson.Get(json, "ip").String()
				port := gjson.Get(json, "port").String()
				path := gjson.Get(json, "path").String()
				parsedTarget := ip + ":" + port + path

				targetsChan <- parsedTarget

				atomic.AddUint64(&targetCtr, 1)
			}(jsonFileLine)
		}
		wg.Wait()
		log.Info().Uint64("targets-num", targetCtr).Msg("Finished parsing targets")
	}()

	return targetsChan
}
