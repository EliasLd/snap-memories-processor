package processor

import (
	"fmt"
	"os/exec"
)

func WriteGPSMetadata(
	file string,
	lat float64,
	lon float64,
) error {

	args := []string{
		fmt.Sprintf("-GPSLatitude=%f", lat),
		fmt.Sprintf("-GPSLongitude=%f", lon),

		"-GPSLatitudeRef=N",
		"-GPSLongitudeRef=E",

		"-overwrite_original",

		file,
	}

	if lat < 0 {

		args[2] = "-GPSLatitudeRef=S"
	}

	if lon < 0 {

		args[3] = "-GPSLongitudeRef=W"
	}

	cmd := exec.Command(
		"exiftool",
		args...,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {

		return fmt.Errorf(
			"exiftool failed: %w\n%s",
			err,
			string(output),
		)
	}

	return nil
}
