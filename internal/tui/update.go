package tui

import (
	"github.com/EliasLd/snap-memories-processor/internal/model"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func Update(
	msg tea.Msg,
	m Model,
) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:

		m.width = msg.Width
		m.height = msg.Height

	case ProgressMsg:

		m.processed = msg.Processed
		m.total = msg.Total

		if m.total > 0 {

			percent := float64(m.processed) / float64(m.total)
			cmd := m.progress.SetPercent(percent)

			return m,
				tea.Batch(
					cmd,
					WaitProgress(
						m.progressChan,
					),
				)

		}

		return m,
			WaitProgress(
				m.progressChan,
			)

	case progress.FrameMsg:

		var cmd tea.Cmd

		pm, cmd := m.progress.Update(msg)

		m.progress = pm.(progress.Model)

		return m, cmd

	case FinishedMsg:

		m.state = StateFinished
		m.summary = msg.summary

		return m, nil

	case ErrorMsg:

		m.lastError = msg.Err
		m.state = StateError

		return m, nil

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

	if m.state == StateError {

		switch msg := msg.(type) {

		case tea.KeyMsg:

			switch msg.String() {

			case "enter":

				m.lastError = nil
				m.state = StateConfig

			case "ctrl+c", "q":

				return m, tea.Quit
			}
		}

		return m, nil
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

				if err := ValidateConfiguration(
					m.inputPath,
					m.gpsEnabled,
				); err != nil {

					m.lastError = err
					m.state = StateError

					return m, nil
				}

				m.state = StateProcessing

				progressChan := make(
					chan model.Progress,
				)

				m.progressChan = progressChan

				return m,
					tea.Batch(
						StartProcessing(
							m.inputPath,
							m.gpsEnabled,
							m.workers,
							progressChan,
						),
						WaitProgress(
							progressChan,
						),
					)
			}
		}
	}

	return m, nil
}
