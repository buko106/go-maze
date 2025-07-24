// Package maze provides maze generation and representation functionality.
// This file implements Unicode rendering for maze output using box-drawing characters.
package maze

import "strings"

// UnicodeRenderer renders mazes using Unicode box-drawing characters.
type UnicodeRenderer struct{}

// Render generates a Unicode representation of the maze using box-drawing characters.
// Uses box-drawing characters for walls, ' ' for paths, '◉' for start, '◎' for goal, and '•' for solution path.
func (r *UnicodeRenderer) Render(m *Maze) string {
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
				sb.WriteRune('◉') // Filled circle with dot for start
			} else if i == m.GoalRow && j == m.GoalCol {
				sb.WriteRune('◎') // Circle with dot for goal
			} else if len(m.SolutionPath) > 0 && solutionSet[currentPos] {
				sb.WriteRune('•') // Bullet for solution path
			} else if cell {
				// Determine appropriate box-drawing character based on connections
				char := r.getBoxDrawingChar(m, i, j)
				sb.WriteRune(char)
			} else {
				sb.WriteRune(' ') // Space for paths
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

// getBoxDrawingChar determines the appropriate box-drawing character for a wall cell
// based on its connections to adjacent wall cells
func (r *UnicodeRenderer) getBoxDrawingChar(m *Maze, row, col int) rune {
	// Check connections in four directions
	up := row > 0 && m.Grid[row-1][col]
	down := row < m.Height-1 && m.Grid[row+1][col]
	left := col > 0 && m.Grid[row][col-1]
	right := col < m.Width-1 && m.Grid[row][col+1]

	// Use a more efficient approach with fewer branches
	return r.selectBoxChar(up, down, left, right)
}

// selectBoxChar returns the box drawing character based on connections using a lookup table
func (r *UnicodeRenderer) selectBoxChar(up, down, left, right bool) rune {
	// Create a lookup table to reduce cyclomatic complexity
	// Index calculation: up*8 + down*4 + left*2 + right*1
	var lookupTable = [16]rune{
		'▪', // 0000: none
		'╶', // 0001: right
		'╴', // 0010: left
		'─', // 0011: left+right
		'╷', // 0100: down
		'┌', // 0101: down+right
		'┐', // 0110: down+left
		'┬', // 0111: down+left+right
		'╵', // 1000: up
		'└', // 1001: up+right
		'┘', // 1010: up+left
		'┴', // 1011: up+left+right
		'│', // 1100: up+down
		'├', // 1101: up+down+right
		'┤', // 1110: up+down+left
		'┼', // 1111: up+down+left+right
	}

	index := 0
	if up {
		index += 8
	}
	if down {
		index += 4
	}
	if left {
		index += 2
	}
	if right {
		index += 1
	}

	return lookupTable[index]
}
