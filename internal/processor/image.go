package processor

func MergeImageOverlay(
	imagePath string,
	overlayPath string,
	outputPath string,
) error {

	return RunFFmpeg(
		"-y",

		"-i",
		imagePath,

		"-i",
		overlayPath,

		"-filter_complex",
		"[1:v][0:v]scale2ref[ov][base];[base][ov]overlay=0:0",

		"-frames:v",
		"1",

		"-update",
		"1",

		outputPath,
	)
}
