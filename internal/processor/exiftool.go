package processor

import "os/exec"

func HasExiftool() bool {

	_, err := exec.LookPath(
		"exiftool",
	)

	return err == nil
}
