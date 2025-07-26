package maze

import (
	"math/rand"
	"strings"
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	analyzer := NewAnalyzer()
	if analyzer == nil {
		t.Error("Expected analyzer instance, got nil")
	}
}

func TestAnalyzeNilMaze(t *testing.T) {
	analyzer := NewAnalyzer()
	analysis, err := analyzer.Analyze(nil)

	if err == nil {
		t.Error("Expected error for nil maze, got nil")
	}
	if analysis != nil {
		t.Error("Expected nil analysis for nil maze, got instance")
	}
}

func TestAnalyzeBasicMetrics(t *testing.T) {
	analyzer := NewAnalyzer()

	// Create a simple test maze
	maze := createTestMaze(5, 5)
	// Generate a basic path for analysis
	dfs := &DFSAlgorithm{}
	rng := rand.New(rand.NewSource(42))
	dfs.Generate(maze, 1, 1, rng)

	analysis, err := analyzer.Analyze(maze)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
		return
	}
	if analysis == nil {
		t.Error("Expected analysis instance, got nil")
		return
	}

	// Verify basic metrics
	expectedTotal := 5 * 5
	if analysis.TotalCells != expectedTotal {
		t.Errorf("Expected total cells %d, got %d", expectedTotal, analysis.TotalCells)
	}

	if analysis.PathCells+analysis.WallCells != analysis.TotalCells {
		t.Error("Path cells + wall cells should equal total cells")
	}

	if analysis.PathDensity < 0 || analysis.PathDensity > 1 {
		t.Errorf("Path density should be between 0 and 1, got %f", analysis.PathDensity)
	}
}

func TestAnalyzeComplexityMetrics(t *testing.T) {
	analyzer := NewAnalyzer()

	// Create and generate a maze
	maze := createTestMaze(7, 7)
	dfs := &DFSAlgorithm{}
	rng := rand.New(rand.NewSource(42))
	dfs.Generate(maze, 1, 1, rng)

	analysis, err := analyzer.Analyze(maze)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify complexity metrics
	totalAnalyzed := analysis.DeadEnds + analysis.Junctions + analysis.Corridors
	if totalAnalyzed != analysis.PathCells {
		t.Errorf("Sum of dead ends (%d) + junctions (%d) + corridors (%d) = %d should equal path cells %d",
			analysis.DeadEnds, analysis.Junctions, analysis.Corridors, totalAnalyzed, analysis.PathCells)
	}

	// Verify ratios are valid
	if analysis.DeadEndRatio < 0 || analysis.DeadEndRatio > 1 {
		t.Errorf("Dead end ratio should be between 0 and 1, got %f", analysis.DeadEndRatio)
	}
	if analysis.JunctionRatio < 0 || analysis.JunctionRatio > 1 {
		t.Errorf("Junction ratio should be between 0 and 1, got %f", analysis.JunctionRatio)
	}
}

func TestAnalyzeSolutionMetrics(t *testing.T) {
	analyzer := NewAnalyzer()

	// Create a maze with known start and goal
	maze := createTestMazeWithStartGoal(7, 7)

	dfs := &DFSAlgorithm{}
	rng := rand.New(rand.NewSource(42))
	dfs.Generate(maze, maze.StartRow, maze.StartCol, rng)

	// Ensure start and goal are paths
	maze.Grid[maze.StartRow][maze.StartCol] = false
	maze.Grid[maze.GoalRow][maze.GoalCol] = false

	analysis, err := analyzer.Analyze(maze)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Solution should exist and be reasonable
	if analysis.SolutionLength <= 0 {
		t.Error("Expected positive solution length")
	}

	if analysis.OptimalRatio < 0 || analysis.OptimalRatio > 1 {
		t.Errorf("Optimal ratio should be between 0 and 1, got %f", analysis.OptimalRatio)
	}
}

func TestAnalyzeDifficultyScoring(t *testing.T) {
	analyzer := NewAnalyzer()

	// Test different algorithms to see difficulty variation
	algorithms := []struct {
		name      string
		algorithm Algorithm
	}{
		{"DFS", &DFSAlgorithm{}},
		{"Hunt-Kill", &HuntKillAlgorithm{}},
		{"Growing-Tree-Easy", NewGrowingTreeWithDifficulty(DifficultyEasy)},
		{"Growing-Tree-Hard", NewGrowingTreeWithDifficulty(DifficultyHard)},
	}

	rng := rand.New(rand.NewSource(42))

	for _, alg := range algorithms {
		t.Run(alg.name, func(t *testing.T) {
			maze := createTestMazeWithStartGoal(11, 11)

			alg.algorithm.Generate(maze, maze.StartRow, maze.StartCol, rng)

			// Ensure start and goal are paths
			maze.Grid[maze.StartRow][maze.StartCol] = false
			maze.Grid[maze.GoalRow][maze.GoalCol] = false

			analysis, err := analyzer.Analyze(maze)
			if err != nil {
				t.Errorf("Expected no error for %s, got: %v", alg.name, err)
			}

			// Difficulty score should be reasonable
			if analysis.DifficultyScore < 0 || analysis.DifficultyScore > 100 {
				t.Errorf("Difficulty score should be between 0 and 100, got %f for %s",
					analysis.DifficultyScore, alg.name)
			}

			// Difficulty level should be assigned
			validLevels := []string{"Trivial", "Easy", "Medium", "Hard", "Extreme"}
			levelValid := false
			for _, level := range validLevels {
				if analysis.DifficultyLevel == level {
					levelValid = true
					break
				}
			}
			if !levelValid {
				t.Errorf("Invalid difficulty level '%s' for %s", analysis.DifficultyLevel, alg.name)
			}
		})
	}
}

func TestAnalyzeCountPathNeighbors(t *testing.T) {
	analyzer := NewAnalyzer()

	// Create a simple maze with known structure
	maze := &Maze{
		Width:  5,
		Height: 5,
		Grid: [][]bool{
			{true, true, true, true, true},
			{true, false, false, false, true},
			{true, false, true, false, true},
			{true, false, false, false, true},
			{true, true, true, true, true},
		},
	}

	// Test corner (dead end)
	count := analyzer.countPathNeighbors(maze, 1, 1)
	if count != 2 {
		t.Errorf("Expected 2 neighbors for corner cell, got %d", count)
	}

	// Test middle corridor
	count = analyzer.countPathNeighbors(maze, 1, 2)
	if count != 2 {
		t.Errorf("Expected 2 neighbors for corridor cell, got %d", count)
	}

	// Test junction
	count = analyzer.countPathNeighbors(maze, 2, 1)
	if count != 2 {
		t.Errorf("Expected 2 neighbors for this specific cell, got %d", count)
	}
}

func TestAnalysisStringOutput(t *testing.T) {
	analyzer := NewAnalyzer()

	// Create and analyze a maze
	maze := createTestMaze(7, 7)
	dfs := &DFSAlgorithm{}
	rng := rand.New(rand.NewSource(42))
	dfs.Generate(maze, 1, 1, rng)

	analysis, err := analyzer.Analyze(maze)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	output := analysis.String()

	// Check that output contains expected sections
	expectedSections := []string{
		"Maze Analysis Report:",
		"Basic Metrics:",
		"Complexity Metrics:",
		"Solution Metrics:",
		"Difficulty Assessment:",
	}

	for _, section := range expectedSections {
		if !strings.Contains(output, section) {
			t.Errorf("Expected output to contain '%s'", section)
		}
	}

	// Check that key metrics are present
	expectedMetrics := []string{
		"Total Cells:",
		"Path Cells:",
		"Dead Ends:",
		"Solution Length:",
		"Difficulty Score:",
		"Difficulty Level:",
	}

	for _, metric := range expectedMetrics {
		if !strings.Contains(output, metric) {
			t.Errorf("Expected output to contain '%s'", metric)
		}
	}
}

func TestAnalyzeDifferentAlgorithmCharacteristics(t *testing.T) {
	analyzer := NewAnalyzer()
	rng := rand.New(rand.NewSource(42)) // Fixed seed for reproducible tests

	// Compare characteristics of different algorithms
	algorithms := map[string]Algorithm{
		"DFS":          &DFSAlgorithm{},
		"Hunt-Kill":    &HuntKillAlgorithm{},
		"Kruskal":      &KruskalAlgorithm{},
		"Wilson":       &WilsonAlgorithm{},
		"Growing-Tree": NewGrowingTreeWithDifficulty(DifficultyMedium),
	}

	results := make(map[string]*MazeAnalysis)

	for name, algorithm := range algorithms {
		maze := createTestMazeWithStartGoal(11, 11)

		algorithm.Generate(maze, maze.StartRow, maze.StartCol, rng)

		// Ensure start and goal are paths (like Generator.Generate does)
		maze.Grid[maze.StartRow][maze.StartCol] = false
		maze.Grid[maze.GoalRow][maze.GoalCol] = false

		analysis, err := analyzer.Analyze(maze)
		if err != nil {
			t.Errorf("Analysis failed for %s: %v", name, err)
			continue
		}

		results[name] = analysis

		t.Logf("%s: Score=%.1f, Level=%s, DeadEnds=%d, Solution=%d",
			name, analysis.DifficultyScore, analysis.DifficultyLevel,
			analysis.DeadEnds, analysis.SolutionLength)
	}

	// All algorithms should produce valid results
	if len(results) != len(algorithms) {
		t.Errorf("Expected %d algorithm results, got %d", len(algorithms), len(results))
	}
}

// Helper function to create a test maze with start and goal positions set
func createTestMazeWithStartGoal(width, height int) *Maze {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
		for j := range grid[i] {
			grid[i][j] = true // Start with all walls
		}
	}
	return &Maze{
		Width:    width,
		Height:   height,
		Grid:     grid,
		StartRow: 1,
		StartCol: 1,
		GoalRow:  height - 2,
		GoalCol:  width - 2,
	}
}
