package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/buko106/go-maze/internal/maze"
)

func main() {
	size := flag.Int("s", 21, "Size of the square maze (must be odd, minimum 5)")
	flag.IntVar(size, "size", 21, "Size of the square maze (must be odd, minimum 5)")
	seed := flag.String("seed", "", "Seed for reproducible maze generation (integer)")
	algorithm := flag.String("a", "dfs", "Algorithm for maze generation (dfs)")
	flag.StringVar(algorithm, "algorithm", "dfs", "Algorithm for maze generation (dfs)")
	flag.Parse()

	// Validate size
	if *size < 5 {
		fmt.Fprintf(os.Stderr, "Error: Size must be at least 5, got %d\n", *size)
		os.Exit(1)
	}
	if *size%2 == 0 {
		fmt.Fprintf(os.Stderr, "Error: Size must be odd, got %d\n", *size)
		os.Exit(1)
	}

	// Validate algorithm
	supportedAlgorithms := maze.GetSupportedAlgorithms()
	algorithmValid := false
	for _, alg := range supportedAlgorithms {
		if *algorithm == alg {
			algorithmValid = true
			break
		}
	}
	if !algorithmValid {
		fmt.Fprintf(os.Stderr, "Error: Unsupported algorithm '%s', supported algorithms: %v\n", *algorithm, supportedAlgorithms)
		os.Exit(1)
	}

	var generator *maze.Generator
	var err error

	if *seed != "" {
		generator, err = maze.NewGeneratorWithSeedAndAlgorithm(*seed, *algorithm)
	} else {
		generator, err = maze.NewGeneratorWithAlgorithm(*algorithm)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating generator: %v\n", err)
		os.Exit(1)
	}

	m := generator.Generate(*size, *size)
	fmt.Print(m.String())
}
