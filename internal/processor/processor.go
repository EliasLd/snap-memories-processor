package processor

import (
	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func ProcessCollection(
	collection []model.Media,
	outputDir string,
	workers int,
) []model.ProcessResult {

	return ProcessAll(
		collection,
		outputDir,
		workers,
	)
}
