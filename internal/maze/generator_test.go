package maze

import (
	"strings"
	"testing"
)

func TestGenerateMaze(t *testing.T) {
	generator := NewGenerator()
	maze := generator.Generate(5, 5)

	if maze.Width != 5 {
		t.Errorf("Expected width 5, got %d", maze.Width)
	}
	if maze.Height != 5 {
		t.Errorf("Expected height 5, got %d", maze.Height)
	}
}

func TestMazeBoundaries(t *testing.T) {
	generator := NewGenerator()
	maze := generator.Generate(5, 5)

	// 外周チェック
	for i := 0; i < maze.Height; i++ {
		if !maze.Grid[i][0] || !maze.Grid[i][maze.Width-1] {
			t.Error("Left or right boundary should be wall")
		}
	}

	for j := 0; j < maze.Width; j++ {
		if !maze.Grid[0][j] || !maze.Grid[maze.Height-1][j] {
			t.Error("Top or bottom boundary should be wall")
		}
	}
}

func TestMazeString(t *testing.T) {
	generator := NewGenerator()
	maze := generator.Generate(5, 5)
	output := maze.String()

	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 5 {
		t.Errorf("Expected 5 lines, got %d", len(lines))
	}

	// Check that output contains expected characters (walls, paths, start, goal)
	if !strings.Contains(output, "#") {
		t.Error("Expected maze to contain wall characters (#)")
	}
	if !strings.Contains(output, "●") {
		t.Error("Expected maze to contain start marker (●)")
	}
	if !strings.Contains(output, "○") {
		t.Error("Expected maze to contain goal marker (○)")
	}
}

// Test that maze has proper path connectivity
func TestMazePathConnectivity(t *testing.T) {
	generator := NewGeneratorWithSeed("123") // Use seed for reproducible testing
	maze := generator.Generate(7, 7)

	// Find all path cells (non-wall cells)
	pathCells := findPathCells(maze)
	if len(pathCells) == 0 {
		t.Error("Maze should have at least one path cell")
		return
	}

	// Test that all path cells are connected
	visited := make(map[[2]int]bool)
	startCell := pathCells[0]
	dfsVisit(maze, startCell[0], startCell[1], visited)

	// Check if all path cells were visited (i.e., all are connected)
	for _, cell := range pathCells {
		if !visited[cell] {
			t.Errorf("Path cell at (%d, %d) is not connected to other paths", cell[0], cell[1])
		}
	}
}

// Helper function to find all path cells
func findPathCells(maze *Maze) [][2]int {
	var pathCells [][2]int
	for i := 0; i < maze.Height; i++ {
		for j := 0; j < maze.Width; j++ {
			if !maze.Grid[i][j] { // false = path
				pathCells = append(pathCells, [2]int{i, j})
			}
		}
	}
	return pathCells
}

// DFS to visit all connected path cells
func dfsVisit(maze *Maze, row, col int, visited map[[2]int]bool) {
	if row < 0 || row >= maze.Height || col < 0 || col >= maze.Width {
		return
	}
	if maze.Grid[row][col] { // wall
		return
	}
	if visited[[2]int{row, col}] {
		return
	}

	visited[[2]int{row, col}] = true

	// Visit neighbors
	dfsVisit(maze, row-1, col, visited) // up
	dfsVisit(maze, row+1, col, visited) // down
	dfsVisit(maze, row, col-1, visited) // left
	dfsVisit(maze, row, col+1, visited) // right
}

// Test that maze displays start and goal markers correctly
func TestMazeStartGoalMarkers(t *testing.T) {
	generator := NewGeneratorWithSeed("123")
	maze := generator.Generate(7, 7)
	output := maze.String()

	lines := strings.Split(strings.TrimSpace(output), "\n")

	// Check start position (top-left path cell) has filled circle ●
	if !strings.Contains(lines[1], "●") {
		t.Error("Expected start marker (●) in maze output")
	}

	// Check goal position (bottom-right path cell) has empty circle ○
	if !strings.Contains(lines[len(lines)-2], "○") {
		t.Error("Expected goal marker (○) in maze output")
	}
}

// Test NewGeneratorWithAlgorithm function
func TestNewGeneratorWithAlgorithm(t *testing.T) {
	// Test valid algorithm
	generator, err := NewGeneratorWithAlgorithm("dfs")
	if err != nil {
		t.Errorf("Expected no error for valid algorithm, got: %v", err)
	}
	if generator == nil {
		t.Error("Expected generator to be created")
	}

	// Test invalid algorithm
	generator, err = NewGeneratorWithAlgorithm("invalid")
	if err == nil {
		t.Error("Expected error for invalid algorithm")
	}
	if generator != nil {
		t.Error("Expected nil generator for invalid algorithm")
	}
}

// Test NewGeneratorWithSeedAndAlgorithm function
func TestNewGeneratorWithSeedAndAlgorithm(t *testing.T) {
	// Test valid algorithm with numeric seed
	generator, err := NewGeneratorWithSeedAndAlgorithm("123", "dfs")
	if err != nil {
		t.Errorf("Expected no error for valid algorithm, got: %v", err)
	}
	if generator == nil {
		t.Error("Expected generator to be created")
	}

	// Test valid algorithm with string seed
	generator, err = NewGeneratorWithSeedAndAlgorithm("test-seed", "dfs")
	if err != nil {
		t.Errorf("Expected no error for string seed, got: %v", err)
	}
	if generator == nil {
		t.Error("Expected generator to be created")
	}

	// Test invalid algorithm
	generator, err = NewGeneratorWithSeedAndAlgorithm("123", "invalid")
	if err == nil {
		t.Error("Expected error for invalid algorithm")
	}
	if generator != nil {
		t.Error("Expected nil generator for invalid algorithm")
	}
}

// Test hashString function indirectly through string seed
func TestHashStringFunctionality(t *testing.T) {
	// Test that string seeds produce consistent results
	generator1 := NewGeneratorWithSeed("test-string")
	generator2 := NewGeneratorWithSeed("test-string")

	maze1 := generator1.Generate(7, 7)
	maze2 := generator2.Generate(7, 7)

	if maze1.String() != maze2.String() {
		t.Error("Same string seed should produce identical mazes")
	}

	// Test different string seeds produce different results
	generator3 := NewGeneratorWithSeed("different-string")
	maze3 := generator3.Generate(7, 7)

	if maze1.String() == maze3.String() {
		t.Error("Different string seeds should produce different mazes")
	}
}

// Test maze string output with solution path
func TestMazeStringWithSolutionPath(t *testing.T) {
	generator := NewGeneratorWithSeed("123")
	maze := generator.Generate(7, 7)

	// Find an actual path using the pathfinder to ensure valid solution path
	solutionPath := FindPath(maze)
	if solutionPath == nil {
		t.Fatal("Could not find path in generated maze")
	}

	maze.SolutionPath = solutionPath

	output := maze.String()

	// Check that solution path markers are displayed
	if !strings.Contains(output, "·") {
		t.Error("Expected maze with solution path to contain solution markers (·)")
	}

	// Check that start and goal markers are still present
	if !strings.Contains(output, "●") {
		t.Error("Expected maze to contain start marker (●)")
	}
	if !strings.Contains(output, "○") {
		t.Error("Expected maze to contain goal marker (○)")
	}

	// Verify that the solution path has the expected number of positions
	if len(maze.SolutionPath) < 2 {
		t.Error("Solution path should have at least start and goal positions")
	}

	// Verify that start and goal are in the solution path
	foundStart := false
	foundGoal := false
	for _, pos := range maze.SolutionPath {
		if pos.Row == maze.StartRow && pos.Col == maze.StartCol {
			foundStart = true
		}
		if pos.Row == maze.GoalRow && pos.Col == maze.GoalCol {
			foundGoal = true
		}
	}
	if !foundStart {
		t.Error("Solution path should include start position")
	}
	if !foundGoal {
		t.Error("Solution path should include goal position")
	}
}

// Test maze string output with empty solution path
func TestMazeStringWithEmptySolutionPath(t *testing.T) {
	generator := NewGeneratorWithSeed("123")
	maze := generator.Generate(7, 7)

	// Explicitly set empty solution path
	maze.SolutionPath = []Position{}

	output := maze.String()

	// Check that no solution path markers are displayed
	if strings.Contains(output, "·") {
		t.Error("Expected maze with empty solution path to not contain solution markers (·)")
	}

	// Check that start and goal markers are still present
	if !strings.Contains(output, "●") {
		t.Error("Expected maze to contain start marker (●)")
	}
	if !strings.Contains(output, "○") {
		t.Error("Expected maze to contain goal marker (○)")
	}
}
