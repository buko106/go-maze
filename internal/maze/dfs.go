package maze

import "math/rand"

// DFSAlgorithm implements maze generation using Depth-First Search
type DFSAlgorithm struct{}

// Generate implements the Algorithm interface using DFS
func (d *DFSAlgorithm) Generate(maze *Maze, startRow, startCol int, rng *rand.Rand) {
	d.generateDFS(maze, startRow, startCol, rng)
}

// generateDFS implements Depth-First Search maze generation algorithm
func (d *DFSAlgorithm) generateDFS(maze *Maze, startRow, startCol int, rng *rand.Rand) {
	// Mark starting cell as path
	maze.Grid[startRow][startCol] = false

	// Define directions: up, right, down, left
	directions := [][2]int{{-2, 0}, {0, 2}, {2, 0}, {0, -2}}

	// Shuffle directions for randomness
	d.shuffleDirections(directions, rng)

	// Try each direction
	for _, dir := range directions {
		newRow := startRow + dir[0]
		newCol := startCol + dir[1]

		// Check if new position is valid and unvisited
		if d.isValidCell(maze, newRow, newCol) && maze.Grid[newRow][newCol] {
			// Remove wall between current and new cell
			wallRow := startRow + dir[0]/2
			wallCol := startCol + dir[1]/2
			maze.Grid[wallRow][wallCol] = false

			// Recursively generate from new cell
			d.generateDFS(maze, newRow, newCol, rng)
		}
	}
}

// shuffleDirections randomizes the order of directions
func (d *DFSAlgorithm) shuffleDirections(directions [][2]int, rng *rand.Rand) {
	for i := len(directions) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		directions[i], directions[j] = directions[j], directions[i]
	}
}

// isValidCell checks if a cell position is within bounds
func (d *DFSAlgorithm) isValidCell(maze *Maze, row, col int) bool {
	return row > 0 && row < maze.Height-1 && col > 0 && col < maze.Width-1
}
