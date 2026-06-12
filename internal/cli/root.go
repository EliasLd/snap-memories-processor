package cli

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/EliasLd/snap-memories-processor/internal/archive"
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

		fmt.Printf("Input:   %s\n", cfg.InputDir)
		fmt.Printf("Output:  %s\n", cfg.OutputDir)
		fmt.Printf("Workers: %d\n\n", cfg.Workers)

		fmt.Printf(
			"\nExtracted %d archive(s)\n\n",
			len(extractions),
		)

		for _, extraction := range extractions {
			fmt.Printf(
				"- %s -> %s\n",
				extraction.ArchiveName,
				extraction.Path,
			)
		}

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
