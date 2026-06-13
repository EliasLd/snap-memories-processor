package processor

import (
	"os"
	"path/filepath"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func ProcessMedia(
	media model.Media,
	outputDir string,
) model.ProcessResult {

	outputFile := filepath.Join(
		outputDir,
		filepath.Base(media.MainPath),
	)

	if err := os.MkdirAll(
		filepath.Dir(outputFile),
		0755,
	); err != nil {

		return model.ProcessResult{
			InputFile:  media.MainPath,
			OutputFile: outputFile,
			Success:    false,
			Error:      err.Error(),
		}
	}

	var err error

	if media.HasOverlay {

		switch media.Metadata.MediaType {

		case "Video":

			err = MergeVideoOverlay(
				media.MainPath,
				media.OverlayPath,
				outputFile,
			)

		case "Image":

			err = MergeImageOverlay(
				media.MainPath,
				media.OverlayPath,
				outputFile,
			)

		default:

			err = CopyMedia(
				media.MainPath,
				outputFile,
			)
		}

	} else {

		err = CopyMedia(
			media.MainPath,
			outputFile,
		)
	}

	if err != nil {

		return model.ProcessResult{
			InputFile:  media.MainPath,
			OutputFile: outputFile,
			Success:    false,
			Error:      err.Error(),
		}
	}

	return model.ProcessResult{
		InputFile:  media.MainPath,
		OutputFile: outputFile,
		Success:    true,
	}
}
