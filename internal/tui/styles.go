package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	selectedStyle = lipgloss.NewStyle().
			Bold(true)

	subtitleStyle = lipgloss.NewStyle().
			Faint(true)

	buttonStyle = lipgloss.NewStyle().
			Padding(0, 2)

	buttonSelectedStyle = lipgloss.NewStyle().
				Bold(true).
				Padding(0, 5)

	titleStyle = lipgloss.NewStyle().
			Bold(true)

	helperStyle = lipgloss.NewStyle().
			Faint(true)

	quoteStyle = lipgloss.NewStyle().
			Italic(true).
			Faint(true)
)
