package tui

func View(m Model) string {

	switch m.state {

	case StateConfig:
		return configView(m)

	case StateProcessing:
		return "Processing..."

	default:
		return ""
	}
}
