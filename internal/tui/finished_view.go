package tui

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)

func finishedView(
	m Model,
) string {

	fullOutputPath, _ := filepath.Abs(m.outputPath)

	content := lipgloss.JoinVertical(
		lipgloss.Center,

		Banner,

		"",

		titleStyle.Render(
			"Processing complete",
		),

		"",

		fmt.Sprintf(
			"Success : %d",
			m.summary.Success,
		),

		fmt.Sprintf(
			"Failed  : %d",
			m.summary.Failed,
		),

		"",
		fullOutputPath,
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
