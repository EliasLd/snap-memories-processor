package tui

type AppState int

const (
	StateConfig AppState = iota
	StateProcessing
	StateFinished
	StateError
)
