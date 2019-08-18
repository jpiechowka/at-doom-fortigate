package doom

import (
	"at-doom-fortigate/config"
	"at-doom-fortigate/files"
	"github.com/rs/zerolog/log"
)

func MainPipeline() {
	log.Info().Msg("Launching pipeline")

	targetsChan := files.ReadRapid7JsonWithTargets(config.Rapid7InputJsonFilePath)
	responsesChan := FireRequestsWithPayloads(targetsChan)
	cleanedAndParsedResponses := CleanAndParseResponses(responsesChan)
	files.WriteResultsToOutputFile(cleanedAndParsedResponses)

	log.Info().
		Str("file", config.OutputFilePath).
		Msg("Pipeline finished! Check output file")
}
