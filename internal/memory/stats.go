package memory

import "github.com/EliasLd/snap-memories-processor/internal/model"

type Stats struct {
	Total       int
	Videos      int
	Images      int
	WithOverlay int
}

func ComputeStats(
	collection []model.Media,
) Stats {

	stats := Stats{
		Total: len(collection),
	}

	for _, media := range collection {

		switch media.Metadata.MediaType {

		case "Video":
			stats.Videos++

		case "Image":
			stats.Images++
		}

		if media.HasOverlay {
			stats.WithOverlay++
		}
	}

	return stats
}
