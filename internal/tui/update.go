package tui

import tea "github.com/charmbracelet/bubbletea"

func Update(
	msg tea.Msg,
	m Model,
) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:

		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "up", "k":

			if m.focus > 0 {
				m.focus--
			}

		case "down", "j":

			if m.focus < FocusCount-1 {
				m.focus++
			}

		case "left", "h":

			if m.focus == FocusWorkers {

				if m.workers > 1 {
					m.workers--
				}
			}

		case "right", "l":

			if m.focus == FocusWorkers {

				if m.workers < 18 {
					m.workers++
				}
			}

		case " ", "enter":

			switch m.focus {

			case FocusGPS:

				m.gpsEnabled = !m.gpsEnabled

			case FocusStart:

				m.state = StateProcessing
			}
		}
	}

	return m, nil
}
