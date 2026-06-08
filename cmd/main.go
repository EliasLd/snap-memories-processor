package main

import (
	"log"

	"github.com/EliasLd/snap-memories-processor/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
