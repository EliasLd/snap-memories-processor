package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func errorView(
	m Model,
) string {

	content := lipgloss.JoinVertical(
		lipgloss.Center,

		ErrorStyle.Render(
			"Unable to start processing",
		),

		"",

		subtitleStyle.Render(
			m.lastError.Error(),
		),

		"",

		quoteStyle.Render(
			"Press Enter to go back",
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
