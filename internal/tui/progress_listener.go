package tui

import (
	"github.com/charmbracelet/bubbletea"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func WaitProgress(
	progress <-chan model.Progress,
) tea.Cmd {

	return func() tea.Msg {

		p, ok := <-progress

		if !ok {

			return nil
		}

		return ProgressMsg{
			Processed: p.Processed,
			Total:     p.Total,
		}
	}
}
