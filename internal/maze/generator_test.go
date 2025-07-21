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
