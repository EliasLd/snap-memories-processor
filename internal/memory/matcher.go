package memory

import (
	"os"
	"time"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func MatchMetadata(
	medias []model.Media,
	metadata map[time.Time]model.Metadata,
) ([]model.Media, int) {

	matched := 0

	for i := range medias {

		info, err := os.Stat(
			medias[i].MainPath,
		)
		if err != nil {
			continue
		}

		timestamp := info.ModTime().UTC()

		meta, ok := metadata[timestamp]
		if !ok {
			continue
		}

		medias[i].Metadata = meta

		matched++
	}

	return medias, matched
}
