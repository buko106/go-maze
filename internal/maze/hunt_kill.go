package maze

import "math/rand"

// HuntKillAlgorithm implements maze generation using Hunt and Kill algorithm
// This algorithm creates more complex mazes with longer paths and more dead ends compared to DFS
type HuntKillAlgorithm struct{}

// Generate implements the Algorithm interface using Hunt and Kill
func (h *HuntKillAlgorithm) Generate(maze *Maze, startRow, startCol int, rng *rand.Rand) {
	h.generateHuntKill(maze, startRow, startCol, rng)
}

// generateHuntKill implements Hunt and Kill maze generation algorithm
func (h *HuntKillAlgorithm) generateHuntKill(maze *Maze, startRow, startCol int, rng *rand.Rand) {
	// Mark starting cell as path
	maze.Grid[startRow][startCol] = false

	// Current position for the "kill" phase
	currentRow, currentCol := startRow, startCol

	for {
		// Kill phase: perform random walk until dead end
		h.killPhase(maze, currentRow, currentCol, rng)

		// Hunt phase: find unvisited cell with visited neighbor
		huntRow, huntCol, found := h.huntPhase(maze, rng)
		if !found {
			break // All cells have been visited
		}

		// Start new kill phase from hunted cell
		currentRow, currentCol = huntRow, huntCol
	}
}

// killPhase performs random walk until reaching a dead end
func (h *HuntKillAlgorithm) killPhase(maze *Maze, startRow, startCol int, rng *rand.Rand) {
	currentRow, currentCol := startRow, startCol

	// Define directions: up, right, down, left
	directions := [][2]int{{-2, 0}, {0, 2}, {2, 0}, {0, -2}}

	for {
		// Get valid unvisited neighbors
		validDirections := h.getValidDirections(maze, currentRow, currentCol, directions)

		if len(validDirections) == 0 {
			// Dead end reached, exit kill phase
			return
		}

		// Choose random direction from valid options
		chosenDir := validDirections[rng.Intn(len(validDirections))]
		newRow := currentRow + chosenDir[0]
		newCol := currentCol + chosenDir[1]

		// Carve path to new cell
		maze.Grid[newRow][newCol] = false

		// Remove wall between current and new cell
		wallRow := currentRow + chosenDir[0]/2
		wallCol := currentCol + chosenDir[1]/2
		maze.Grid[wallRow][wallCol] = false

		// Move to new cell
		currentRow, currentCol = newRow, newCol
	}
}

// huntPhase scans grid to find unvisited cell with visited neighbor
// Returns position of found cell and whether one was found
func (h *HuntKillAlgorithm) huntPhase(maze *Maze, rng *rand.Rand) (int, int, bool) {
	// Scan grid systematically
	for row := 1; row < maze.Height-1; row += 2 {
		for col := 1; col < maze.Width-1; col += 2 {
			// Check if cell is unvisited (wall)
			if maze.Grid[row][col] {
				// Check if it has at least one visited neighbor
				visitedNeighbors := h.getVisitedNeighbors(maze, row, col)
				if len(visitedNeighbors) > 0 {
					// Choose random visited neighbor to connect to
					chosenNeighbor := visitedNeighbors[rng.Intn(len(visitedNeighbors))]

					// Carve path to hunted cell
					maze.Grid[row][col] = false

					// Remove wall between hunted cell and chosen neighbor
					wallRow := row + chosenNeighbor[0]/2
					wallCol := col + chosenNeighbor[1]/2
					maze.Grid[wallRow][wallCol] = false

					return row, col, true
				}
			}
		}
	}

	return 0, 0, false // No unvisited cells with visited neighbors found
}

// getValidDirections returns list of directions that lead to valid unvisited cells
func (h *HuntKillAlgorithm) getValidDirections(maze *Maze, row, col int, directions [][2]int) [][2]int {
	var validDirections [][2]int

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		// Check if new position is valid and unvisited
		if h.isValidCell(maze, newRow, newCol) && maze.Grid[newRow][newCol] {
			validDirections = append(validDirections, dir)
		}
	}

	return validDirections
}

// getVisitedNeighbors returns list of directions to visited neighboring cells
func (h *HuntKillAlgorithm) getVisitedNeighbors(maze *Maze, row, col int) [][2]int {
	var visitedNeighbors [][2]int
	directions := [][2]int{{-2, 0}, {0, 2}, {2, 0}, {0, -2}}

	for _, dir := range directions {
		neighborRow := row + dir[0]
		neighborCol := col + dir[1]

		// Check if neighbor is valid and visited (path)
		if h.isValidCell(maze, neighborRow, neighborCol) && !maze.Grid[neighborRow][neighborCol] {
			visitedNeighbors = append(visitedNeighbors, dir)
		}
	}

	return visitedNeighbors
}

// isValidCell checks if a cell position is within bounds
func (h *HuntKillAlgorithm) isValidCell(maze *Maze, row, col int) bool {
	return row > 0 && row < maze.Height-1 && col > 0 && col < maze.Width-1
}
