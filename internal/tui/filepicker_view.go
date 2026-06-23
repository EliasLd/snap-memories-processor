package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func filepickerView(
	m Model,
) string {

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		"",
		titleStyle.Render(
			"Select directory containing snap exports.",
		),
		subtitleStyle.Render(
			"Current: ",
			m.filepicker.CurrentDirectory,
		),
		"",
		m.filepicker.View(),
	)

	helper := helperStyle.Render(
		"• q quit •",
	)

	view := lipgloss.JoinVertical(
		lipgloss.Center,
		content,
		helper,
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		view,
	)
}
