package memory

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func BuildCollection(
	extractions []model.Extraction,
) ([]model.Media, error) {

	var jsonPath string

	for _, extraction := range extractions {

		candidate := filepath.Join(
			extraction.Path,
			"json",
			"memories_history.json",
		)

		if _, err := os.Stat(candidate); err == nil {
			jsonPath = candidate
			break
		}
	}

	if jsonPath == "" {
		return nil, fmt.Errorf(
			"memories_history.json not found",
		)
	}

	metadata, err := LoadMetadata(
		jsonPath,
	)
	if err != nil {
		return nil, err
	}

	var allMedia []model.Media

	for _, extraction := range extractions {

		memoriesDir := filepath.Join(
			extraction.Path,
			"memories",
		)

		medias, err := ScanMemories(
			memoriesDir,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"scan memories from %s: %w",
				extraction.ArchiveName,
				err,
			)
		}

		allMedia = append(
			allMedia,
			medias...,
		)
	}

	allMedia, matches := MatchMetadata(
		allMedia,
		metadata,
	)

	if matches != len(allMedia) {

		return nil, fmt.Errorf(
			"matched %d/%d media",
			matches,
			len(allMedia),
		)
	}

	return allMedia, nil
}
