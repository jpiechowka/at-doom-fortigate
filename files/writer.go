package files

import (
	"at-doom-fortigate/config"
	"at-doom-fortigate/logging"
	"bufio"
	"errors"
	"github.com/rs/zerolog/log"
	"os"
)

func WriteResultsToOutputFile(parsedCleanedResponsesChan <-chan []byte) {
	log.Info().Str("file", config.OutputFilePath).Msg("Creating output file")
	fileToWrite, fileCreateError := os.Create(config.OutputFilePath)
	logging.CheckFatalError(fileCreateError, "Error when creating output file")

	defer func() {
		log.Info().
			Str("file", config.OutputFilePath).
			Msg("Writing done - closing file")
		fileCloseError := fileToWrite.Close()
		logging.CheckFatalError(fileCloseError, "Error when closing file in defer block")
	}()

	bufioWriter := bufio.NewWriter(fileToWrite)

	for response := range parsedCleanedResponsesChan {
		bytesWritten, bytesWriteError := bufioWriter.Write(response)
		logging.CheckFatalError(bytesWriteError, "Error when writing bytes to the buffer")

		var bytesWrittenMismatchError error
		if bytesWritten != len(response) {
			bytesWrittenMismatchError = errors.New("The length of bytes written: " + string(bytesWritten) + " is not " +
				"equal to the length of bytes that needed to be written: " + string(len(response)))
		}
		logging.CheckFatalError(bytesWrittenMismatchError, "Error when writing bytes to the buffer")

		bufferedWriterFlushError := bufioWriter.Flush()
		logging.CheckFatalError(bufferedWriterFlushError, "Error when flushing buffer")
	}

	log.Info().
		Str("file", config.OutputFilePath).
		Msg("Syncing output file")

	fileSyncError := fileToWrite.Sync()
	logging.CheckFatalError(fileSyncError, "Error when syncing output file")

}
