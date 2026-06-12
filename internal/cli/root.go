package cli

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/EliasLd/snap-memories-processor/internal/archive"
	"github.com/EliasLd/snap-memories-processor/internal/memory"
)

type Config struct {
	InputDir  string
	OutputDir string
	Workers   int
}

var cfg Config

var rootCmd = &cobra.Command{
	Use:   "snapmemories",
	Short: "Snapchat memories exporter",
}

var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Process Snapchat exports",
	RunE: func(cmd *cobra.Command, args []string) error {

		archives, err := archive.Discover(cfg.InputDir)
		if err != nil {
			return err
		}

		tmpDir := "./tmp/extracted"

		extractions, err := archive.ExtractAll(
			archives,
			tmpDir,
		)
		if err != nil {
			return err
		}

		var allMetadata int
		var allMedia int
		var totalMatches int

		for _, extraction := range extractions {

			jsonPath := filepath.Join(
				extraction.Path,
				"json",
				"memories_history.json",
			)

			metadata, err := memory.LoadMetadata(
				jsonPath,
			)
			if err != nil {
				return err
			}

			allMetadata += len(metadata)

			memoriesDir := filepath.Join(
				extraction.Path,
				"memories",
			)

			medias, err := memory.ScanMemories(
				memoriesDir,
			)
			if err != nil {
				return err
			}

			allMedia += len(medias)

			_, matches := memory.MatchMetadata(
				medias,
				metadata,
			)

			totalMatches += matches
		}

		fmt.Println()

		fmt.Printf(
			"Metadata loaded : %d\n",
			allMetadata,
		)

		fmt.Printf(
			"Media found     : %d\n",
			allMedia,
		)

		fmt.Printf(
			"Matches         : %d\n",
			totalMatches,
		)

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	processCmd.Flags().StringVarP(
		&cfg.InputDir,
		"input",
		"i",
		"",
		"Directory containing Snapchat exports",
	)

	processCmd.Flags().StringVarP(
		&cfg.OutputDir,
		"output",
		"o",
		"./output",
		"Output directory",
	)

	processCmd.Flags().IntVarP(
		&cfg.Workers,
		"workers",
		"w",
		runtime.NumCPU(),
		"Number of workers",
	)

	processCmd.MarkFlagRequired("input")

	rootCmd.AddCommand(processCmd)
}
