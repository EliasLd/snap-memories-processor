package processor

import (
	"sync"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func ProcessAll(
	medias []model.Media,
	outputDir string,
	workers int,
) []model.ProcessResult {

	jobs := make(
		chan model.Media,
		workers*2,
	)

	results := make(
		chan model.ProcessResult,
		workers*2,
	)

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			for media := range jobs {

				results <- ProcessMedia(
					media,
					outputDir,
				)
			}
		}()
	}

	go func() {

		for _, media := range medias {
			jobs <- media
		}

		close(jobs)

		wg.Wait()

		close(results)
	}()

	processResults := make(
		[]model.ProcessResult,
		0,
		len(medias),
	)

	for result := range results {

		processResults = append(
			processResults,
			result,
		)
	}

	return processResults
}
