package archive

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func Discover(dir string) ([]model.Archive, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read input directory: %w", err)
	}

	var archives []model.Archive

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) != ".zip" {
			continue
		}

		archives = append(archives, model.Archive{
			Name: entry.Name(),
			Path: filepath.Join(dir, entry.Name()),
		})
	}

	return archives, nil
}
