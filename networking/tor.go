package networking

import (
	"at-doom-fortigate/config"
	"github.com/rs/zerolog/log"
	"strings"
)

func IsTorConnected() bool {
	log.Info().Msg("Checking TOR connection")

	torCheckResponse := GetRequestThroughTor(config.TorCheckURL)

	if torCheckResponse.HttpStatusCode >= 400 {
		log.Error().
			Int("response-status-code", torCheckResponse.HttpStatusCode).
			Msg("HTTP status code is higher than or equal 400. TOR is not connected")
		return false
	}

	log.Info().
		Str("networking-check-url", config.TorCheckURL).
		Msg("Checking TOR response body from TOR check")

	if strings.Contains(strings.ToLower(string(torCheckResponse.ResponseBody)), strings.ToLower(config.YouAreUsingTorString)) {
		log.Info().Msg("TOR is connected! Ready to rock!")
		return true
	}

	log.Error().Msg("TOR is not connected")
	return false
}
