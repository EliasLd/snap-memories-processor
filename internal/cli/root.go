package cli

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/EliasLd/snap-memories-processor/internal/archive"
	"github.com/EliasLd/snap-memories-processor/internal/memory"
	"github.com/EliasLd/snap-memories-processor/internal/processor"
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
		collection, err := memory.BuildCollection(
			extractions,
		)
		if err != nil {
			return err
		}

		stats := memory.ComputeStats(
			collection,
		)

		fmt.Println()

		fmt.Printf("Total media  : %d\n", stats.Total)
		fmt.Printf("Videos       : %d\n", stats.Videos)
		fmt.Printf("Images       : %d\n", stats.Images)
		fmt.Printf("With overlay : %d\n", stats.WithOverlay)

		results := processor.ProcessCollection(
			collection,
			cfg.OutputDir,
			cfg.Workers,
		)

		success := processor.CountSuccess(
			results,
		)

		failures := processor.CountFailures(
			results,
		)

		fmt.Println()

		fmt.Printf(
			"Processed     : %d\n",
			success,
		)

		fmt.Printf(
			"Failed        : %d\n",
			failures,
		)

		if failures > 0 {

			err := processor.WriteErrorLog(
				results,
				cfg.OutputDir,
			)
			if err != nil {
				return err
			}

			fmt.Printf(
				"\nErrors written to %s/errors.log\n",
				cfg.OutputDir,
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
