package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func processingView(
	m Model,
) string {

	content := lipgloss.JoinVertical(
		lipgloss.Center,

		Banner,

		"",

		titleStyle.Render(
			"Processing memories...",
		),

		"",

		fmt.Sprintf(
			"%d / %d",
			m.processed,
			m.total,
		),
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
