package tui

type Focus int

const (
	FocusGPS Focus = iota
	FocusWorkers
	FocusStart
)

const FocusCount = 3
