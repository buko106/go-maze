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
	maze := generator.Generate(3, 3)
	output := maze.String()

	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(lines))
	}

	for _, line := range lines {
		if len(line) != 3 {
			t.Errorf("Expected line length 3, got %d", len(line))
		}
	}
}
