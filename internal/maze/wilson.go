package maze

import "math/rand"

// WilsonAlgorithm implements maze generation using Wilson's algorithm
type WilsonAlgorithm struct{}

// Generate implements the Algorithm interface using Wilson's algorithm
func (w *WilsonAlgorithm) Generate(maze *Maze, startRow, startCol int, rng *rand.Rand) {
	w.generateWilson(maze, startRow, startCol, rng)
}

// generateWilson implements Wilson's algorithm using loop-erased random walks
func (w *WilsonAlgorithm) generateWilson(maze *Maze, startRow, startCol int, rng *rand.Rand) {
	// Track which cells are part of the maze
	inMaze := make([][]bool, maze.Height)
	for i := range inMaze {
		inMaze[i] = make([]bool, maze.Width)
	}

	// Add the starting cell to the maze
	maze.Grid[startRow][startCol] = false
	inMaze[startRow][startCol] = true

	// Get all cells that can be part of paths (odd coordinates)
	cells := w.getAllPathCells(maze)

	// Remove the starting cell from the list
	cells = w.removeCellFromList(cells, startRow, startCol)

	// Process each remaining cell
	for len(cells) > 0 {
		// Pick a random cell not yet in the maze
		cellIndex := rng.Intn(len(cells))
		currentRow, currentCol := cells[cellIndex][0], cells[cellIndex][1]

		// Skip if this cell is already in the maze
		if inMaze[currentRow][currentCol] {
			cells = append(cells[:cellIndex], cells[cellIndex+1:]...)
			continue
		}

		// Perform loop-erased random walk
		path := w.loopErasedRandomWalk(maze, currentRow, currentCol, inMaze, rng)

		// Add the entire path to the maze
		for i, pos := range path {
			maze.Grid[pos[0]][pos[1]] = false
			inMaze[pos[0]][pos[1]] = true

			// Connect to previous cell in path (remove wall between them)
			if i > 0 {
				prev := path[i-1]
				wallRow := (pos[0] + prev[0]) / 2
				wallCol := (pos[1] + prev[1]) / 2
				maze.Grid[wallRow][wallCol] = false
			}
		}

		// Remove processed cells from the list
		cells = w.removeProcessedCells(cells, inMaze)
	}
}

// loopErasedRandomWalk performs a random walk with loop erasure
func (w *WilsonAlgorithm) loopErasedRandomWalk(maze *Maze, startRow, startCol int, inMaze [][]bool, rng *rand.Rand) [][2]int {
	// Track the current path and position in path for loop detection
	path := make([][2]int, 0)
	pathIndex := make(map[[2]int]int) // maps position to index in path

	currentRow, currentCol := startRow, startCol

	for {
		pos := [2]int{currentRow, currentCol}

		// If we've reached a cell that's already in the maze, we're done
		if inMaze[currentRow][currentCol] {
			path = append(path, pos)
			break
		}

		// Check if we've created a loop
		if index, exists := pathIndex[pos]; exists {
			// Erase the loop by truncating the path
			path = path[:index]
			// Update pathIndex to remove erased positions
			newPathIndex := make(map[[2]int]int)
			for i, p := range path {
				newPathIndex[p] = i
			}
			pathIndex = newPathIndex
		}

		// Add current position to path
		pathIndex[pos] = len(path)
		path = append(path, pos)

		// Choose a random direction
		directions := [][2]int{{-2, 0}, {0, 2}, {2, 0}, {0, -2}} // up, right, down, left
		validDirections := make([][2]int, 0)

		for _, dir := range directions {
			newRow := currentRow + dir[0]
			newCol := currentCol + dir[1]
			if w.isValidCell(maze, newRow, newCol) {
				validDirections = append(validDirections, dir)
			}
		}

		// Move in a random valid direction
		if len(validDirections) > 0 {
			dir := validDirections[rng.Intn(len(validDirections))]
			currentRow += dir[0]
			currentCol += dir[1]
		}
	}

	return path
}

// getAllPathCells returns all cells that can be part of paths (odd coordinates)
func (w *WilsonAlgorithm) getAllPathCells(maze *Maze) [][2]int {
	cells := make([][2]int, 0)
	for row := 1; row < maze.Height; row += 2 {
		for col := 1; col < maze.Width; col += 2 {
			cells = append(cells, [2]int{row, col})
		}
	}
	return cells
}

// removeCellFromList removes a specific cell from the cell list
func (w *WilsonAlgorithm) removeCellFromList(cells [][2]int, row, col int) [][2]int {
	for i, cell := range cells {
		if cell[0] == row && cell[1] == col {
			return append(cells[:i], cells[i+1:]...)
		}
	}
	return cells
}

// removeProcessedCells removes cells that are now in the maze from the cell list
func (w *WilsonAlgorithm) removeProcessedCells(cells [][2]int, inMaze [][]bool) [][2]int {
	remaining := make([][2]int, 0)
	for _, cell := range cells {
		if !inMaze[cell[0]][cell[1]] {
			remaining = append(remaining, cell)
		}
	}
	return remaining
}

// isValidCell checks if a cell position is within bounds
func (w *WilsonAlgorithm) isValidCell(maze *Maze, row, col int) bool {
	return row > 0 && row < maze.Height-1 && col > 0 && col < maze.Width-1
}
