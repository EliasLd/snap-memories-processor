package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Update(m Model, msg tea.Msg) (Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:

		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	}

	return m, nil
}
