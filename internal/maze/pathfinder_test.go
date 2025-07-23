package maze

import (
	"testing"
)

// TestFindPathBasic tests basic pathfinding functionality
func TestFindPathBasic(t *testing.T) {
	// Create a simple 5x5 maze with a clear path
	maze := &Maze{
		Width:    5,
		Height:   5,
		StartRow: 1,
		StartCol: 1,
		GoalRow:  3,
		GoalCol:  3,
		Grid: [][]bool{
			{true, true, true, true, true},    // All walls on top
			{true, false, false, false, true}, // Start at (1,1)
			{true, false, true, false, true},  // Wall in middle
			{true, false, false, false, true}, // Goal at (3,3)
			{true, true, true, true, true},    // All walls on bottom
		},
	}

	path := FindPath(maze)

	if path == nil {
		t.Fatal("Expected path to be found, got nil")
	}

	// Check that path starts at start position
	if path[0].Row != maze.StartRow || path[0].Col != maze.StartCol {
		t.Errorf("Path should start at (%d,%d), got (%d,%d)",
			maze.StartRow, maze.StartCol, path[0].Row, path[0].Col)
	}

	// Check that path ends at goal position
	lastIdx := len(path) - 1
	if path[lastIdx].Row != maze.GoalRow || path[lastIdx].Col != maze.GoalCol {
		t.Errorf("Path should end at (%d,%d), got (%d,%d)",
			maze.GoalRow, maze.GoalCol, path[lastIdx].Row, path[lastIdx].Col)
	}

	// Check that all positions in path are valid (not walls)
	for i, pos := range path {
		if maze.Grid[pos.Row][pos.Col] {
			t.Errorf("Position %d in path (%d,%d) is a wall", i, pos.Row, pos.Col)
		}
	}
}

// TestFindPathNoPath tests when no path exists
func TestFindPathNoPath(t *testing.T) {
	// Create a maze where goal is blocked
	maze := &Maze{
		Width:    5,
		Height:   5,
		StartRow: 1,
		StartCol: 1,
		GoalRow:  3,
		GoalCol:  3,
		Grid: [][]bool{
			{true, true, true, true, true},
			{true, false, false, false, true},
			{true, false, true, false, true},
			{true, false, true, true, true}, // Goal is blocked
			{true, true, true, true, true},
		},
	}

	path := FindPath(maze)

	if path != nil {
		t.Error("Expected no path to be found, but got a path")
	}
}

// TestFindPathStartBlocked tests when start position is blocked
func TestFindPathStartBlocked(t *testing.T) {
	maze := &Maze{
		Width:    5,
		Height:   5,
		StartRow: 1,
		StartCol: 1,
		GoalRow:  3,
		GoalCol:  3,
		Grid: [][]bool{
			{true, true, true, true, true},
			{true, true, false, false, true}, // Start is blocked
			{true, false, true, false, true},
			{true, false, false, false, true},
			{true, true, true, true, true},
		},
	}

	path := FindPath(maze)

	if path != nil {
		t.Error("Expected no path when start is blocked, but got a path")
	}
}

// TestFindPathGoalBlocked tests when goal position is blocked
func TestFindPathGoalBlocked(t *testing.T) {
	maze := &Maze{
		Width:    5,
		Height:   5,
		StartRow: 1,
		StartCol: 1,
		GoalRow:  3,
		GoalCol:  3,
		Grid: [][]bool{
			{true, true, true, true, true},
			{true, false, false, false, true},
			{true, false, true, false, true},
			{true, false, false, true, true}, // Goal is blocked
			{true, true, true, true, true},
		},
	}

	path := FindPath(maze)

	if path != nil {
		t.Error("Expected no path when goal is blocked, but got a path")
	}
}

// TestFindPathSameStartGoal tests when start and goal are the same
func TestFindPathSameStartGoal(t *testing.T) {
	maze := &Maze{
		Width:    3,
		Height:   3,
		StartRow: 1,
		StartCol: 1,
		GoalRow:  1,
		GoalCol:  1,
		Grid: [][]bool{
			{true, true, true},
			{true, false, true},
			{true, true, true},
		},
	}

	path := FindPath(maze)

	if path == nil {
		t.Fatal("Expected path to be found when start equals goal")
	}

	if len(path) != 1 {
		t.Errorf("Expected path length 1 when start equals goal, got %d", len(path))
	}

	if path[0].Row != 1 || path[0].Col != 1 {
		t.Errorf("Expected path to contain only (1,1), got (%d,%d)", path[0].Row, path[0].Col)
	}
}

// TestFindPathConnectivity tests pathfinding on generated mazes
func TestFindPathConnectivity(t *testing.T) {
	generator := NewGeneratorWithSeed("test123")
	maze := generator.Generate(11, 11)

	path := FindPath(maze)

	if path == nil {
		t.Fatal("Expected path to be found in generated maze")
	}

	// Verify path connectivity (each step should be adjacent)
	for i := 1; i < len(path); i++ {
		prev := path[i-1]
		curr := path[i]

		// Check if adjacent (Manhattan distance should be 1)
		distance := abs(curr.Row-prev.Row) + abs(curr.Col-prev.Col)
		if distance != 1 {
			t.Errorf("Path positions %d and %d are not adjacent: (%d,%d) -> (%d,%d)",
				i-1, i, prev.Row, prev.Col, curr.Row, curr.Col)
		}
	}
}

// Helper function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
