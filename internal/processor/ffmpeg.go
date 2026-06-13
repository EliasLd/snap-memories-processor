package processor

import (
	"fmt"
	"os/exec"
)

func RunFFmpeg(args ...string) error {

	cmd := exec.Command(
		"ffmpeg",
		args...,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {

		return fmt.Errorf(
			"ffmpeg failed: %w\n%s",
			err,
			string(output),
		)
	}

	return nil
}
