package maze

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewAlgorithm(t *testing.T) {
	// Test valid DFS algorithm
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

	// Test valid Kruskal algorithm
	algorithm, err = NewAlgorithm("kruskal")
	if err != nil {
		t.Errorf("Expected no error for valid algorithm 'kruskal', got: %v", err)
	}
	if algorithm == nil {
		t.Error("Expected algorithm instance, got nil")
	}
	if _, ok := algorithm.(*KruskalAlgorithm); !ok {
		t.Error("Expected KruskalAlgorithm instance")
	}

	// Test valid Wilson algorithm
	algorithm, err = NewAlgorithm("wilson")
	if err != nil {
		t.Errorf("Expected no error for valid algorithm 'wilson', got: %v", err)
	}
	if algorithm == nil {
		t.Error("Expected algorithm instance, got nil")
	}
	if _, ok := algorithm.(*WilsonAlgorithm); !ok {
		t.Error("Expected WilsonAlgorithm instance")
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
	if len(algorithms) < 3 {
		t.Error("Expected at least three supported algorithms")
	}

	// Check that all algorithms are supported
	dfsSupported := false
	kruskalSupported := false
	wilsonSupported := false
	for _, alg := range algorithms {
		if alg == "dfs" {
			dfsSupported = true
		}
		if alg == "kruskal" {
			kruskalSupported = true
		}
		if alg == "wilson" {
			wilsonSupported = true
		}
	}
	if !dfsSupported {
		t.Error("Expected 'dfs' to be in supported algorithms")
	}
	if !kruskalSupported {
		t.Error("Expected 'kruskal' to be in supported algorithms")
	}
	if !wilsonSupported {
		t.Error("Expected 'wilson' to be in supported algorithms")
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

func TestKruskalAlgorithmGenerate(t *testing.T) {
	kruskal := &KruskalAlgorithm{}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a test maze
	width, height := 5, 5
	maze := createTestMaze(width, height)

	// Generate maze using Kruskal's algorithm
	kruskal.Generate(maze, 1, 1, rng)

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

func TestKruskalReproducibility(t *testing.T) {
	kruskal := &KruskalAlgorithm{}
	seed := int64(12345)

	// Generate first maze
	rng1 := rand.New(rand.NewSource(seed))
	maze1 := createTestMaze(5, 5)
	kruskal.Generate(maze1, 1, 1, rng1)

	// Generate second maze with same seed
	rng2 := rand.New(rand.NewSource(seed))
	maze2 := createTestMaze(5, 5)
	kruskal.Generate(maze2, 1, 1, rng2)

	// Compare mazes
	for i := 0; i < maze1.Height; i++ {
		for j := 0; j < maze1.Width; j++ {
			if maze1.Grid[i][j] != maze2.Grid[i][j] {
				t.Errorf("Mazes differ at position (%d, %d)", i, j)
			}
		}
	}
}

func TestKruskalConnectivity(t *testing.T) {
	kruskal := &KruskalAlgorithm{}
	rng := rand.New(rand.NewSource(42))

	// Create a larger test maze
	width, height := 9, 9
	maze := createTestMaze(width, height)

	// Generate maze using Kruskal's algorithm
	kruskal.Generate(maze, 1, 1, rng)

	// Verify connectivity using flood fill from start position
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	// Flood fill from position (1,1)
	floodFill(maze, visited, 1, 1)

	// Count reachable path cells
	reachableCount := 0
	totalPathCount := 0
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			if !maze.Grid[i][j] { // It's a path
				totalPathCount++
				if visited[i][j] {
					reachableCount++
				}
			}
		}
	}

	// All path cells should be reachable
	if reachableCount != totalPathCount {
		t.Errorf("Not all paths are connected: reachable=%d, total=%d", reachableCount, totalPathCount)
	}
}

// Helper function for flood fill connectivity test
func floodFill(maze *Maze, visited [][]bool, row, col int) {
	if row < 0 || row >= maze.Height || col < 0 || col >= maze.Width {
		return
	}
	if visited[row][col] || maze.Grid[row][col] {
		return
	}

	visited[row][col] = true

	// Recursively fill adjacent cells
	floodFill(maze, visited, row-1, col) // up
	floodFill(maze, visited, row+1, col) // down
	floodFill(maze, visited, row, col-1) // left
	floodFill(maze, visited, row, col+1) // right
}

func TestWilsonAlgorithmGenerate(t *testing.T) {
	wilson := &WilsonAlgorithm{}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a test maze
	width, height := 7, 7
	maze := createTestMaze(width, height)

	// Generate maze using Wilson's algorithm
	wilson.Generate(maze, 1, 1, rng)

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

func TestWilsonReproducibility(t *testing.T) {
	wilson := &WilsonAlgorithm{}
	seed := int64(12345)

	// Generate first maze
	rng1 := rand.New(rand.NewSource(seed))
	maze1 := createTestMaze(7, 7)
	wilson.Generate(maze1, 1, 1, rng1)

	// Generate second maze with same seed
	rng2 := rand.New(rand.NewSource(seed))
	maze2 := createTestMaze(7, 7)
	wilson.Generate(maze2, 1, 1, rng2)

	// Compare mazes
	for i := 0; i < maze1.Height; i++ {
		for j := 0; j < maze1.Width; j++ {
			if maze1.Grid[i][j] != maze2.Grid[i][j] {
				t.Errorf("Mazes differ at position (%d, %d)", i, j)
			}
		}
	}
}

func TestWilsonConnectivity(t *testing.T) {
	wilson := &WilsonAlgorithm{}
	rng := rand.New(rand.NewSource(42))

	// Create a larger test maze
	width, height := 9, 9
	maze := createTestMaze(width, height)

	// Generate maze using Wilson's algorithm
	wilson.Generate(maze, 1, 1, rng)

	// Verify connectivity using flood fill from start position
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	// Flood fill from position (1,1)
	floodFill(maze, visited, 1, 1)

	// Count reachable path cells
	reachableCount := 0
	totalPathCount := 0
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			if !maze.Grid[i][j] { // It's a path
				totalPathCount++
				if visited[i][j] {
					reachableCount++
				}
			}
		}
	}

	// All path cells should be reachable
	if reachableCount != totalPathCount {
		t.Errorf("Not all paths are connected: reachable=%d, total=%d", reachableCount, totalPathCount)
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
