package memory

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

type memoriesHistory struct {
	SavedMedia []savedMedia `json:"Saved Media"`
}

type savedMedia struct {
	Date      string `json:"Date"`
	MediaType string `json:"Media Type"`
	Location  string `json:"Location"`
}

func LoadMetadata(
	jsonPath string,
) (map[time.Time]model.Metadata, error) {

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	var history memoriesHistory

	if err := json.Unmarshal(data, &history); err != nil {
		return nil, err
	}

	metadata := make(map[time.Time]model.Metadata)

	for _, entry := range history.SavedMedia {

		timestamp, err := time.Parse(
			"2006-01-02 15:04:05 MST",
			entry.Date,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"parse timestamp %s: %w",
				entry.Date,
				err,
			)
		}

		lat, lon, err := parseCoordinates(
			entry.Location,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"parse coordinates: %w",
				err,
			)
		}

		metadata[timestamp.UTC()] = model.Metadata{
			Date: timestamp.UTC(),

			Latitude:  lat,
			Longitude: lon,

			MediaType: entry.MediaType,
		}
	}

	return metadata, nil
}

func parseCoordinates(
	location string,
) (float64, float64, error) {

	const prefix = "Latitude, Longitude: "

	location = strings.TrimPrefix(
		location,
		prefix,
	)

	parts := strings.Split(location, ",")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf(
			"invalid location format",
		)
	}

	lat, err := strconv.ParseFloat(
		strings.TrimSpace(parts[0]),
		64,
	)
	if err != nil {
		return 0, 0, err
	}

	lon, err := strconv.ParseFloat(
		strings.TrimSpace(parts[1]),
		64,
	)
	if err != nil {
		return 0, 0, err
	}

	return lat, lon, nil
}
