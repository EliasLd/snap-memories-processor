package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width  int
	height int

	state AppState

	focus Focus

	gpsEnabled bool
	workers    int
}

func InitialModel() Model {
	return Model{
		state:      StateConfig,
		focus:      FocusGPS,
		gpsEnabled: false,
		workers:    16,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return Update(msg, m)
}

func (m Model) View() string {
	return View(m)
}
