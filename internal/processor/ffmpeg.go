package processor

import (
	"fmt"
	"os/exec"
)

func HasFFmpeg() bool {

	_, err := exec.LookPath(
		"ffmpeg",
	)

	return err == nil
}

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
