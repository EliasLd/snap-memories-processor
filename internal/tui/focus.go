package tui

type Focus int

const (
	FocusInput Focus = iota
	FocusGPS
	FocusWorkers
	FocusStart

	FocusCount
)
