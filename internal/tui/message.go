package tui

type ProgressMsg struct {
	Processed int
	Total     int
}

type FinishedMsg struct {
	summary Summary
}

type ErrorMsg struct {
	Err error
}
