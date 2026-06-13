package model

type ProcessResult struct {
	InputFile  string
	OutputFile string

	Success bool
	Error   string
}
