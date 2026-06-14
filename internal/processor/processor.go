package processor

import (
	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func ProcessCollection(
	collection []model.Media,
	outputDir string,
	workers int,
	writeGPS bool,
	progress chan<- model.Progress,
) []model.ProcessResult {

	return ProcessAll(
		collection,
		outputDir,
		workers,
		writeGPS,
		progress,
	)
}
