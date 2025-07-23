package maze

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewAlgorithm(t *testing.T) {
	// Test valid algorithm
	algorithm, err := NewAlgorithm("dfs")
	if err != nil {
		t.Errorf("Expected no error for valid algorithm 'dfs', got: %v", err)
	}
	if algorithm == nil {
		t.Error("Expected algorithm instance, got nil")
	}
	if _, ok := algorithm.(*DFSAlgorithm); !ok {
		t.Error("Expected DFSAlgorithm instance")
	}

	// Test invalid algorithm
	algorithm, err = NewAlgorithm("invalid")
	if err == nil {
		t.Error("Expected error for invalid algorithm, got nil")
	}
	if algorithm != nil {
		t.Error("Expected nil algorithm for invalid name, got instance")
	}
}

func TestGetSupportedAlgorithms(t *testing.T) {
	algorithms := GetSupportedAlgorithms()
	if len(algorithms) == 0 {
		t.Error("Expected at least one supported algorithm")
	}

	// Check that DFS is supported
	dfsSupported := false
	for _, alg := range algorithms {
		if alg == "dfs" {
			dfsSupported = true
			break
		}
	}
	if !dfsSupported {
		t.Error("Expected 'dfs' to be in supported algorithms")
	}
}

func TestDFSAlgorithmGenerate(t *testing.T) {
	dfs := &DFSAlgorithm{}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a test maze
	width, height := 5, 5
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
		for j := range grid[i] {
			grid[i][j] = true // Start with all walls
		}
	}

	maze := &Maze{
		Width:  width,
		Height: height,
		Grid:   grid,
	}

	// Generate maze using DFS
	dfs.Generate(maze, 1, 1, rng)

	// Verify that the starting position is a path
	if maze.Grid[1][1] != false {
		t.Error("Expected starting position (1,1) to be a path, got wall")
	}

	// Verify that some paths were created (not all walls)
	pathCount := 0
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			if !maze.Grid[i][j] {
				pathCount++
			}
		}
	}
	if pathCount == 0 {
		t.Error("Expected some paths to be created, got all walls")
	}

	// Verify boundaries are still walls
	for i := 0; i < height; i++ {
		if !maze.Grid[i][0] || !maze.Grid[i][width-1] {
			t.Error("Expected boundaries to remain walls")
		}
	}
	for j := 0; j < width; j++ {
		if !maze.Grid[0][j] || !maze.Grid[height-1][j] {
			t.Error("Expected boundaries to remain walls")
		}
	}
}

func TestDFSReproducibility(t *testing.T) {
	dfs := &DFSAlgorithm{}
	seed := int64(12345)

	// Generate first maze
	rng1 := rand.New(rand.NewSource(seed))
	maze1 := createTestMaze(5, 5)
	dfs.Generate(maze1, 1, 1, rng1)

	// Generate second maze with same seed
	rng2 := rand.New(rand.NewSource(seed))
	maze2 := createTestMaze(5, 5)
	dfs.Generate(maze2, 1, 1, rng2)

	// Compare mazes
	for i := 0; i < maze1.Height; i++ {
		for j := 0; j < maze1.Width; j++ {
			if maze1.Grid[i][j] != maze2.Grid[i][j] {
				t.Errorf("Mazes differ at position (%d, %d)", i, j)
			}
		}
	}
}

// Helper function to create a test maze with all walls
func createTestMaze(width, height int) *Maze {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
		for j := range grid[i] {
			grid[i][j] = true // Start with all walls
		}
	}
	return &Maze{
		Width:  width,
		Height: height,
		Grid:   grid,
	}
}
