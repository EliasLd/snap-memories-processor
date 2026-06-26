package tui

import (
	"errors"

	"github.com/EliasLd/snap-memories-processor/internal/archive"
	"github.com/EliasLd/snap-memories-processor/internal/processor"
)

func ValidateConfiguration(
	inputDir string,
	gps bool,
) error {

	if !processor.HasFFmpeg() {

		return errors.New(
			"FFmpeg was not found. Please install FFmpeg before using the processor.",
		)
	}

	if gps && !processor.HasExiftool() {

		return errors.New(
			"GPS preservation requires ExifTool to be installed.",
		)
	}

	_, err := archive.Discover(
		inputDir,
	)
	if err != nil {
		return err
	}

	return nil
}
