package processor

import (
	"fmt"
	"os/exec"
)

func WriteVideoGPSMetadata(
	file string,
	lat float64,
	lon float64,
) error {

	coords := fmt.Sprintf(
		"%f %f",
		lat,
		lon,
	)

	cmd := exec.Command(
		"exiftool",

		"-overwrite_original",

		fmt.Sprintf(
			"-GPSCoordinates=%s",
			coords,
		),

		file,
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
