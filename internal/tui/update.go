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

	}

	if m.state == StateFilePicker {

		switch msg := msg.(type) {

		case tea.KeyMsg:

			switch msg.String() {

			case "q":
				m.state = StateConfig

			}

		}

		var cmd tea.Cmd

		m.filepicker, cmd = m.filepicker.Update(msg)

		if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {

			m.inputPath = path
			m.state = StateConfig
		}

		return m, cmd
	}

	// Config state
	switch msg := msg.(type) {

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

			case FocusInput:

				m.state = StateFilePicker

				cmd := m.resetFilepicker()

				return m, tea.Batch(
					cmd,
					func() tea.Msg {
						return tea.WindowSizeMsg{
							Width:  m.width,
							Height: m.height,
						}
					},
				)

			case FocusGPS:

				m.gpsEnabled = !m.gpsEnabled

			case FocusStart:

				m.state = StateProcessing
			}
		}
	}

	return m, nil
}
