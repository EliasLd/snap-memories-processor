package tui

import (
	"os"

	"github.com/EliasLd/snap-memories-processor/internal/model"
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width  int
	height int

	state AppState

	focus Focus

	gpsEnabled bool
	workers    int

	inputPath  string
	outputPath string

	filepicker filepicker.Model

	progress progress.Model

	progressChan chan model.Progress

	processed int
	total     int

	summary Summary
}

func (m *Model) resetFilepicker() tea.Cmd {

	fp := filepicker.New()

	fp.CurrentDirectory, _ = os.Getwd()

	fp.ShowHidden = false

	fp.DirAllowed = true
	fp.FileAllowed = false

	m.filepicker = fp

	return m.filepicker.Init()
}

func InitialModel() Model {
	fp := filepicker.New()

	fp.CurrentDirectory, _ = os.Getwd()

	fp.ShowHidden = false

	fp.DirAllowed = true
	fp.FileAllowed = false

	inputDirPath, _ := os.Getwd()

	p := progress.New(
		progress.WithDefaultGradient(),
	)

	return Model{
		state:      StateConfig,
		focus:      FocusInput,
		gpsEnabled: false,
		workers:    16,
		filepicker: fp,
		inputPath:  inputDirPath,
		outputPath: "./output",
		progress:   p,
	}
}

func (m Model) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return Update(msg, m)
}

func (m Model) View() string {
	return View(m)
}
