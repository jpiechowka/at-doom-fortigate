package main

import (
	"at-doom-fortigate/doom"
	"at-doom-fortigate/networking"
	"github.com/rs/zerolog/log"
)

func main() {
	if !networking.IsTorConnected() {
		log.Fatal().Msg("Have you lost your mind? Real hax0rs cannot operate without TOR")
	}

	// Main logic
	doom.MainPipeline()
}
