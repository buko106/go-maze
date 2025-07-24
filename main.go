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
	algorithm := flag.String("a", "dfs", "Algorithm for maze generation (dfs, kruskal, wilson)")
	flag.StringVar(algorithm, "algorithm", "dfs", "Algorithm for maze generation (dfs, kruskal, wilson)")
	format := flag.String("f", "ascii", "Output format (ascii, unicode, json)")
	flag.StringVar(format, "format", "ascii", "Output format (ascii, unicode, json)")
	solution := flag.Bool("solution", false, "Display the solution path from start to goal")
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

	// Validate format
	supportedFormats := maze.GetSupportedFormats()
	formatValid := false
	for _, fmt := range supportedFormats {
		if *format == fmt {
			formatValid = true
			break
		}
	}
	if !formatValid {
		fmt.Fprintf(os.Stderr, "Error: Unsupported format '%s', supported formats: %v\n", *format, supportedFormats)
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

	// If solution flag is set, compute and display the solution path
	if *solution {
		solutionPath := maze.FindPath(m)
		if solutionPath != nil {
			m.SolutionPath = solutionPath
		}
	}

	// Create renderer and output maze
	renderer, err := maze.NewRenderer(*format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating renderer: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(renderer.Render(m))
}
