// Package maze provides maze generation and representation functionality.
// This file implements JSON rendering for maze output.
package maze

import (
	"encoding/json"
)

// JSONRenderer renders mazes as JSON for programmatic use.
type JSONRenderer struct{}

// JSON represents the JSON structure for maze output.
type JSON struct {
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Grid         [][]bool   `json:"grid"`
	Start        Position   `json:"start"`
	Goal         Position   `json:"goal"`
	SolutionPath []Position `json:"solution_path,omitempty"`
}

// Render generates a JSON representation of the maze.
// Returns a formatted JSON string with maze structure and metadata.
func (r *JSONRenderer) Render(m *Maze) string {
	mazeJSON := JSON{
		Width:  m.Width,
		Height: m.Height,
		Grid:   m.Grid,
		Start: Position{
			Row: m.StartRow,
			Col: m.StartCol,
		},
		Goal: Position{
			Row: m.GoalRow,
			Col: m.GoalCol,
		},
		SolutionPath: m.SolutionPath,
	}

	jsonBytes, err := json.MarshalIndent(mazeJSON, "", "  ")
	if err != nil {
		return "{\"error\": \"failed to marshal maze to JSON\"}"
	}

	return string(jsonBytes)
}
