package model

import "time"

type Metadata struct {
	Date time.Time

	Latitude  float64
	Longitude float64

	MediaType string
}

type Media struct {
	MainPath string

	OverlayPath string

	HasOverlay bool

	Metadata Metadata
}

type Job struct {
	Media     Media
	OutputDir string
	WriteGPS  bool
}

type Result struct {
	File string
	Err  error
}
