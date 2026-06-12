package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func Extract(
	archive model.Archive,
	outputDir string,
) (model.Extraction, error) {

	name := strings.TrimSuffix(
		archive.Name,
		filepath.Ext(archive.Name),
	)

	destDir := filepath.Join(outputDir, name)

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return model.Extraction{}, err
	}

	reader, err := zip.OpenReader(archive.Path)
	if err != nil {
		return model.Extraction{}, err
	}
	defer reader.Close()

	for _, file := range reader.File {

		target := filepath.Join(destDir, file.Name)

		// Avoid Zip Slip
		if !strings.HasPrefix(
			target,
			filepath.Clean(destDir)+string(os.PathSeparator),
		) {
			return model.Extraction{}, fmt.Errorf(
				"illegal file path: %s",
				file.Name,
			)
		}

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(target, 0755); err != nil {
				return model.Extraction{}, err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return model.Extraction{}, err
		}

		src, err := file.Open()
		if err != nil {
			return model.Extraction{}, err
		}

		dst, err := os.Create(target)
		if err != nil {
			src.Close()
			return model.Extraction{}, err
		}

		_, err = io.Copy(dst, src)

		dst.Close()
		src.Close()

		if err != nil {
			return model.Extraction{}, err
		}
	}

	return model.Extraction{
		ArchiveName: archive.Name,
		Path:        destDir,
	}, nil
}

func ExtractAll(
	archives []model.Archive,
	outputDir string,
) ([]model.Extraction, error) {

	var extractions []model.Extraction

	for _, archive := range archives {

		extraction, err := Extract(
			archive,
			outputDir,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"extract %s: %w",
				archive.Name,
				err,
			)
		}

		extractions = append(
			extractions,
			extraction,
		)
	}

	return extractions, nil
}
