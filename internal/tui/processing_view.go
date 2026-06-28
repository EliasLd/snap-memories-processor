package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func processingView(
	m Model,
) string {

	title := titleStyle.Render(
		"Processing Memories",
	)

	var status string

	switch m.phase {

	case PhaseExtracting:

		status = "Extracting Snapchat archives..."

	case PhaseBuildingCollection:

		status = "Reading Snapchat metadata..."

	case PhaseProcessing:

		status = "Processing memories..."
	}

	status = subtitleStyle.Render(
		status,
	)

	bar := m.progress.View()

	content := []string{
		Banner,

		"",
		title,

		"",
		status,

		"",
		bar,
	}

	if m.phase == PhaseProcessing {

		percent := 0.0

		if m.total > 0 {

			percent =
				float64(m.processed) /
					float64(m.total) * 100
		}

		info := subtitleStyle.Render(
			fmt.Sprintf(
				"%d / %d files (%.1f%%)",
				m.processed,
				m.total,
				percent,
			),
		)

		content = append(
			content,
			"",
			info,
		)
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			content...,
		),
	)
}
