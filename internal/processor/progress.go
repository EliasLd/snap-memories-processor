package processor

import (
	"fmt"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func RenderProgress(
	progress <-chan model.Progress,
) {

	const width = 30

	for p := range progress {

		percent := float64(
			p.Processed,
		) / float64(
			p.Total,
		)

		filled := int(
			percent * width,
		)

		fmt.Printf(
			"\r[%s%s] %d/%d (%d%%)",
			repeat("#", filled),
			repeat("-", width-filled),
			p.Processed,
			p.Total,
			int(percent*100),
		)
	}

	fmt.Println()
}

func repeat(
	s string,
	count int,
) string {

	result := ""

	for i := 0; i < count; i++ {
		result += s
	}

	return result
}
