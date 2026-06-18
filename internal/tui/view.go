package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func View(m Model) string {

	const quote string = "Memories are important, let's keep them."

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		titleStyle.Render(Banner),
		"",
		quoteStyle.Render(quote),
	)

	helper := helperStyle.Render(
		"q / esc : quit",
	)

	bodyHeight := max(
		0,
		m.height-2,
	)

	main := lipgloss.Place(
		m.width,
		bodyHeight,

		lipgloss.Center,
		lipgloss.Center,

		content,
	)

	return main + "\n" + helper
}
