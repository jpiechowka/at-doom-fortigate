package logging

import (
	"github.com/rs/zerolog/log"
)

func CheckFatalError(errorToCheck error, errorMessageToLog string) {
	if errorToCheck != nil {
		log.Fatal().Err(errorToCheck).Msg(errorMessageToLog)
	}
}
