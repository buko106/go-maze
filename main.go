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
	algorithm := flag.String("a", "dfs", "Algorithm for maze generation (dfs, kruskal, wilson, hunt-kill, growing-tree)")
	flag.StringVar(algorithm, "algorithm", "dfs", "Algorithm for maze generation (dfs, kruskal, wilson, hunt-kill, growing-tree)")
	difficulty := flag.String("d", "medium", "Difficulty level for growing-tree algorithm (easy, medium, hard, extreme)")
	flag.StringVar(difficulty, "difficulty", "medium", "Difficulty level for growing-tree algorithm (easy, medium, hard, extreme)")
	format := flag.String("f", "ascii", "Output format (ascii, unicode, json)")
	flag.StringVar(format, "format", "ascii", "Output format (ascii, unicode, json)")
	solution := flag.Bool("solution", false, "Display the solution path from start to goal")
	analyze := flag.Bool("analyze", false, "Display detailed maze analysis and complexity metrics")
	flag.Parse()

	// Validate all input parameters
	validateSize(*size)
	validateAlgorithm(*algorithm)
	validateDifficulty(*difficulty)
	validateFormat(*format)

	var generator *maze.Generator
	var err error

	if *seed != "" {
		generator, err = maze.NewGeneratorWithSeedAlgorithmAndDifficulty(*seed, *algorithm, *difficulty)
	} else {
		generator, err = maze.NewGeneratorWithAlgorithmAndDifficulty(*algorithm, *difficulty)
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

	// If analyze flag is set, perform analysis
	var analysis *maze.MazeAnalysis
	if *analyze {
		analyzer := maze.NewAnalyzer()
		var err error
		analysis, err = analyzer.Analyze(m)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error analyzing maze: %v\n", err)
		}
	}

	// Create renderer and output maze
	renderer, err := maze.NewRenderer(*format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating renderer: %v\n", err)
		os.Exit(1)
	}

	// Handle different output formats
	if *format == "json" && *analyze && analysis != nil {
		// For JSON format with analysis, use special renderer method
		if jsonRenderer, ok := renderer.(*maze.JSONRenderer); ok {
			fmt.Print(jsonRenderer.RenderWithAnalysis(m, analysis))
		} else {
			fmt.Print(renderer.Render(m))
		}
	} else {
		// For non-JSON or no analysis, use standard output
		if *analyze && analysis != nil {
			fmt.Println(analysis.String())
			fmt.Println() // Add blank line before maze output
		}
		fmt.Print(renderer.Render(m))
	}
}

// validateSize checks if the size parameter is valid
func validateSize(size int) {
	if size < 5 {
		fmt.Fprintf(os.Stderr, "Error: Size must be at least 5, got %d\n", size)
		os.Exit(1)
	}
	if size%2 == 0 {
		fmt.Fprintf(os.Stderr, "Error: Size must be odd, got %d\n", size)
		os.Exit(1)
	}
}

// validateAlgorithm checks if the algorithm parameter is supported
func validateAlgorithm(algorithm string) {
	supportedAlgorithms := maze.GetSupportedAlgorithms()
	for _, alg := range supportedAlgorithms {
		if algorithm == alg {
			return
		}
	}
	fmt.Fprintf(os.Stderr, "Error: Unsupported algorithm '%s', supported algorithms: %v\n", algorithm, supportedAlgorithms)
	os.Exit(1)
}

// validateDifficulty checks if the difficulty parameter is supported
func validateDifficulty(difficulty string) {
	supportedDifficulties := maze.GetSupportedDifficulties()
	for _, diff := range supportedDifficulties {
		if difficulty == diff {
			return
		}
	}
	fmt.Fprintf(os.Stderr, "Error: Unsupported difficulty '%s', supported difficulties: %v\n", difficulty, supportedDifficulties)
	os.Exit(1)
}

// validateFormat checks if the format parameter is supported
func validateFormat(format string) {
	supportedFormats := maze.GetSupportedFormats()
	for _, fmt := range supportedFormats {
		if format == fmt {
			return
		}
	}
	fmt.Fprintf(os.Stderr, "Error: Unsupported format '%s', supported formats: %v\n", format, supportedFormats)
	os.Exit(1)
}
