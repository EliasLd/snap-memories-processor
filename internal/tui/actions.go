package tui

import (
	"github.com/charmbracelet/bubbletea"

	"github.com/EliasLd/snap-memories-processor/internal/archive"
	"github.com/EliasLd/snap-memories-processor/internal/memory"
	"github.com/EliasLd/snap-memories-processor/internal/model"
	"github.com/EliasLd/snap-memories-processor/internal/processor"
)

func StartProcessing(
	inputDir string,
	gps bool,
	workers int,
	progress chan<- model.Progress,
) tea.Cmd {

	return func() tea.Msg {

		archives, err := archive.Discover(
			inputDir,
		)
		if err != nil {
			return ErrorMsg{Err: err}
		}

		extractions, err := archive.ExtractAll(
			archives,
			"./tmp/extracted",
		)
		if err != nil {
			return ErrorMsg{Err: err}
		}

		collection, err := memory.BuildCollection(
			extractions,
		)
		if err != nil {
			return ErrorMsg{Err: err}
		}

		results := processor.ProcessAll(
			collection,
			"./output",
			workers,
			gps,
			progress,
		)

		success := processor.CountSuccess(
			results,
		)

		failed := processor.CountFailures(
			results,
		)

		return FinishedMsg{
			Summary{
				success,
				failed,
			},
		}
	}
}
