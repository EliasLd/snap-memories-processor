package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width  int
	height int

	state AppState
}

func InitialModel() Model {
	return Model{
		state: StateConfig,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return Update(m, msg)
}

func (m Model) View() string {
	return View(m)
}
