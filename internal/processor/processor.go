package processor

import (
	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func ProcessCollection(
	collection []model.Media,
	outputDir string,
	workers int,
	progress chan<- model.Progress,
) []model.ProcessResult {

	return ProcessAll(
		collection,
		outputDir,
		workers,
		progress,
	)
}
