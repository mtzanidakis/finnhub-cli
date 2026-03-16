package main

import (
	"fmt"
	"os"

	"github.com/mtzanidakis/finnhub-cli/cmd"
)

func main() {
	if len(os.Args) < 2 {
		cmd.Usage()
		os.Exit(1)
	}

	if err := cmd.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
