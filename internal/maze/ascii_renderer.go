// Package maze provides maze generation and representation functionality.
// This file implements ASCII rendering for maze output.
package maze

import "strings"

// ASCIIRenderer renders mazes using standard ASCII characters.
type ASCIIRenderer struct{}

// Render generates an ASCII representation of the maze.
// Uses '#' for walls, ' ' for paths, '●' for start, '○' for goal, and '·' for solution path.
func (r *ASCIIRenderer) Render(m *Maze) string {
	var sb strings.Builder

	// Create a set of solution positions for quick lookup
	solutionSet := make(map[Position]bool)
	for _, pos := range m.SolutionPath {
		solutionSet[pos] = true
	}

	for i, row := range m.Grid {
		for j, cell := range row {
			currentPos := Position{Row: i, Col: j}

			if i == m.StartRow && j == m.StartCol {
				sb.WriteRune('●') // Filled circle for start
			} else if i == m.GoalRow && j == m.GoalCol {
				sb.WriteRune('○') // Empty circle for goal
			} else if len(m.SolutionPath) > 0 && solutionSet[currentPos] {
				sb.WriteRune('·') // Solution path marker
			} else if cell {
				sb.WriteRune('#')
			} else {
				sb.WriteRune(' ')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
