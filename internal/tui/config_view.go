package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func configView(
	m Model,
) string {

	input := fmt.Sprintf(
		"Input directory: %s",
		m.inputPath,
	)

	if m.focus == FocusInput {
		input = "> " + input
		input = selectedStyle.Render(input)

	} else {
		input = subtitleStyle.Render(input)
	}

	gps := "[ ] Preserve GPS metadata"

	if m.gpsEnabled {
		gps = "[x] Preserve GPS metadata"
	}

	if m.focus == FocusGPS {
		gps = "> " + gps
		gps = selectedStyle.Render(gps)
	} else {
		gps = subtitleStyle.Render(gps)
	}

	workers := fmt.Sprintf(
		"Workers: %d",
		m.workers,
	)

	if m.focus == FocusWorkers {
		workers = "> " + workers
		workers = selectedStyle.Render(workers)
	} else {
		workers = subtitleStyle.Render(workers)
	}

	start := "[ Start Process ]"

	if m.focus == FocusStart {
		start = "> " + start
		start = buttonSelectedStyle.Render(start)
	} else {
		start = buttonStyle.Render(start)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Center,

		Banner,

		"",
		quoteStyle.Render(
			"Memories are important, let's keep them.",
		),
		"",
		input,
		"",
		gps,
		"",
		workers,
		"",
		start,
	)

	helper := helperStyle.Render(
		"↑/k up • ↓/j down • ←/→ workers • enter select • q quit",
	)

	view := lipgloss.JoinVertical(
		lipgloss.Center,
		content,
		"",
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
