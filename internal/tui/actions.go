package tui

import (
	"github.com/charmbracelet/bubbletea"

	"github.com/EliasLd/snap-memories-processor/internal/archive"
	"github.com/EliasLd/snap-memories-processor/internal/memory"
	"github.com/EliasLd/snap-memories-processor/internal/model"
	"github.com/EliasLd/snap-memories-processor/internal/processor"
)

var program *tea.Program

func SetProgram(
	p *tea.Program,
) {
	program = p
}

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

		program.Send(
			PhaseMsg{
				Phase: PhaseExtracting,
			},
		)

		extractions, err := archive.ExtractAll(
			archives,
			"./tmp/extracted",
		)
		if err != nil {
			return ErrorMsg{Err: err}
		}

		program.Send(
			PhaseMsg{
				Phase: PhaseBuildingCollection,
			},
		)

		collection, err := memory.BuildCollection(
			extractions,
		)
		if err != nil {
			return ErrorMsg{Err: err}
		}

		program.Send(
			PhaseMsg{
				Phase: PhaseProcessing,
			},
		)

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
