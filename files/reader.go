package files

import (
	"at-doom-fortigate/config"
	"at-doom-fortigate/logging"
	"bufio"
	"github.com/rs/zerolog/log"
	"os"
)

func readFileLines(filePath string) <-chan string {
	fileLinesChan := make(chan string, config.FileLinesChanBufferSize)

	go func() {
		defer close(fileLinesChan)

		log.Info().
			Str("file", filePath).
			Msg("Opening file for reading")
		file, fileOpenError := os.Open(filePath)

		defer func() {
			log.Info().
				Str("file", filePath).
				Msg("Reading done - closing file")
			fileCloseError := file.Close()
			logging.CheckFatalError(fileCloseError, "Error when closing file in defer block")
		}()

		logging.CheckFatalError(fileOpenError, "Error when opening file for reading")

		// Setup bufio scanner
		// Use custom buffer size (bigger than usual) to avoid token too large logging
		scannerBuffer := make([]byte, config.BufioScannerBufferSize)
		scanner := bufio.NewScanner(file)
		scanner.Buffer(scannerBuffer, 0)

		// Read file line by line and send to channel
		for scanner.Scan() {
			fileLinesChan <- scanner.Text()
		}

		logging.CheckFatalError(scanner.Err(), "Error reading file")
	}()

	return fileLinesChan
}
