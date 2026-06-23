package tui

type AppState int

const (
	StateConfig AppState = iota
	StateFilePicker
	StateProcessing
	StateFinished
	StateError
)
