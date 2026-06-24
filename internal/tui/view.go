package tui

func View(m Model) string {

	switch m.state {

	case StateConfig:
		return configView(m)

	case StateFilePicker:
		return filepickerView(m)

	case StateProcessing:
		return processingView(m)

	case StateFinished:
		return finishedView(m)

	default:
		return ""
	}
}
