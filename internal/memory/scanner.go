package memory

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func ScanMemories(
	memoriesDir string,
) ([]model.Media, error) {

	entries, err := os.ReadDir(memoriesDir)
	if err != nil {
		return nil, err
	}

	overlays := make(map[string]string)

	for _, entry := range entries {

		name := entry.Name()

		if !strings.Contains(name, "-overlay.") {
			continue
		}

		key := overlayKey(name)

		overlays[key] = filepath.Join(
			memoriesDir,
			name,
		)
	}

	var medias []model.Media

	for _, entry := range entries {

		name := entry.Name()

		if entry.IsDir() {
			continue
		}

		if name == "memories.html" {
			continue
		}

		if strings.Contains(name, "-overlay.") {
			continue
		}

		if !strings.Contains(name, "-main.") {
			continue
		}

		media := model.Media{
			MainPath: filepath.Join(
				memoriesDir,
				name,
			),
		}

		key := overlayKey(name)

		if overlayPath, ok := overlays[key]; ok {

			media.HasOverlay = true
			media.OverlayPath = overlayPath
		}

		medias = append(
			medias,
			media,
		)
	}

	return medias, nil
}

func overlayKey(
	filename string,
) string {

	filename = strings.Replace(
		filename,
		"-main",
		"",
		1,
	)

	filename = strings.Replace(
		filename,
		"-overlay",
		"",
		1,
	)

	ext := filepath.Ext(filename)

	return strings.TrimSuffix(
		filename,
		ext,
	)
}
