package tui

type ProcessingPhase int

const (
	PhaseExtracting ProcessingPhase = iota
	PhaseBuildingCollection
	PhaseProcessing
)
