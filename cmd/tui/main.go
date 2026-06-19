package main

import (
	"fmt"
	"github.com/EliasLd/snap-memories-processor/internal/tui"
	"os"
)

func main() {
	if err := tui.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
