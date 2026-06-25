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

	bar := m.progress.View()

	content := lipgloss.JoinVertical(
		lipgloss.Center,

		Banner,

		"",
		title,

		"",
		bar,

		"",
		info,
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
