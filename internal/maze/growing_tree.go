package maze

import "math/rand"

// DifficultyLevel represents the difficulty setting for maze generation
type DifficultyLevel string

const (
	// DifficultyEasy creates mazes with more branching and shorter paths
	DifficultyEasy DifficultyLevel = "easy"
	// DifficultyMedium creates balanced mazes with moderate complexity
	DifficultyMedium DifficultyLevel = "medium"
	// DifficultyHard creates mazes with longer winding paths and more dead ends
	DifficultyHard DifficultyLevel = "hard"
	// DifficultyExtreme creates very complex mazes with maximum path length and complexity
	DifficultyExtreme DifficultyLevel = "extreme"
)

// SelectionStrategy defines how cells are selected from the active list
type SelectionStrategy struct {
	NewestWeight float64 // Weight for selecting newest cell (DFS-like behavior)
	OldestWeight float64 // Weight for selecting oldest cell (BFS-like behavior)
	RandomWeight float64 // Weight for random selection
}

// GrowingTreeAlgorithm implements maze generation using Growing Tree algorithm
// This algorithm allows fine-grained control over maze characteristics and difficulty
type GrowingTreeAlgorithm struct {
	Difficulty DifficultyLevel
	Strategy   SelectionStrategy
}

// NewGrowingTreeWithDifficulty creates a Growing Tree algorithm with predefined difficulty settings
func NewGrowingTreeWithDifficulty(difficulty DifficultyLevel) *GrowingTreeAlgorithm {
	strategy := getStrategyForDifficulty(difficulty)
	return &GrowingTreeAlgorithm{
		Difficulty: difficulty,
		Strategy:   strategy,
	}
}

// Generate implements the Algorithm interface using Growing Tree
func (g *GrowingTreeAlgorithm) Generate(maze *Maze, startRow, startCol int, rng *rand.Rand) {
	g.generateGrowingTree(maze, startRow, startCol, rng)
}

// generateGrowingTree implements Growing Tree maze generation algorithm
func (g *GrowingTreeAlgorithm) generateGrowingTree(maze *Maze, startRow, startCol int, rng *rand.Rand) {
	// Mark starting cell as path
	maze.Grid[startRow][startCol] = false

	// Active list contains cells that can potentially expand
	activeList := []Position{{Row: startRow, Col: startCol}}

	// Continue until no more cells can expand
	for len(activeList) > 0 {
		// Select cell from active list based on strategy
		selectedIndex := g.selectCellIndex(activeList, rng)
		selectedCell := activeList[selectedIndex]

		// Try to expand from selected cell
		if newCell := g.expandFromCell(maze, selectedCell, rng); newCell != nil {
			// Cell expanded successfully, add new cell to active list
			activeList = append(activeList, *newCell)
		} else {
			// Cell cannot expand, remove from active list
			activeList = append(activeList[:selectedIndex], activeList[selectedIndex+1:]...)
		}
	}
}

// selectCellIndex chooses which cell from the active list to process based on strategy
func (g *GrowingTreeAlgorithm) selectCellIndex(activeList []Position, rng *rand.Rand) int {
	if len(activeList) == 1 {
		return 0
	}

	// Calculate total weight
	totalWeight := g.Strategy.NewestWeight + g.Strategy.OldestWeight + g.Strategy.RandomWeight
	if totalWeight == 0 {
		return rng.Intn(len(activeList)) // Fallback to random
	}

	// Generate random value in range [0, totalWeight)
	randomValue := rng.Float64() * totalWeight

	// Determine selection method based on weights
	if randomValue < g.Strategy.NewestWeight {
		// Select newest cell (DFS-like)
		return len(activeList) - 1
	} else if randomValue < g.Strategy.NewestWeight+g.Strategy.OldestWeight {
		// Select oldest cell (BFS-like)
		return 0
	}
	// Select random cell
	return rng.Intn(len(activeList))
}

// expandFromCell attempts to carve a path from the given cell to an unvisited neighbor
// Returns the position of the newly carved cell, or nil if no expansion was possible
func (g *GrowingTreeAlgorithm) expandFromCell(maze *Maze, cell Position, rng *rand.Rand) *Position {
	// Define directions: up, right, down, left
	directions := [][2]int{{-2, 0}, {0, 2}, {2, 0}, {0, -2}}

	// Get valid unvisited neighbors
	validDirections := []int{}
	for i, dir := range directions {
		newRow := cell.Row + dir[0]
		newCol := cell.Col + dir[1]

		// Check if new position is valid and unvisited
		if g.isValidCell(maze, newRow, newCol) && maze.Grid[newRow][newCol] {
			validDirections = append(validDirections, i)
		}
	}

	if len(validDirections) == 0 {
		return nil // No expansion possible
	}

	// Choose random direction from valid options
	chosenDirIndex := validDirections[rng.Intn(len(validDirections))]
	chosenDir := directions[chosenDirIndex]

	newRow := cell.Row + chosenDir[0]
	newCol := cell.Col + chosenDir[1]

	// Carve path to new cell
	maze.Grid[newRow][newCol] = false

	// Remove wall between current and new cell
	wallRow := cell.Row + chosenDir[0]/2
	wallCol := cell.Col + chosenDir[1]/2
	maze.Grid[wallRow][wallCol] = false

	return &Position{Row: newRow, Col: newCol}
}

// isValidCell checks if a cell position is within bounds
func (g *GrowingTreeAlgorithm) isValidCell(maze *Maze, row, col int) bool {
	return row > 0 && row < maze.Height-1 && col > 0 && col < maze.Width-1
}

// getStrategyForDifficulty returns predefined selection strategies for different difficulty levels
func getStrategyForDifficulty(difficulty DifficultyLevel) SelectionStrategy {
	switch difficulty {
	case DifficultyEasy:
		// More BFS-like behavior creates wider, more branching mazes (easier)
		return SelectionStrategy{
			NewestWeight: 0.2, // Less DFS behavior
			OldestWeight: 0.6, // More BFS behavior
			RandomWeight: 0.2, // Some randomness
		}
	case DifficultyMedium:
		// Balanced approach
		return SelectionStrategy{
			NewestWeight: 0.4, // Moderate DFS behavior
			OldestWeight: 0.3, // Moderate BFS behavior
			RandomWeight: 0.3, // Good mix of randomness
		}
	case DifficultyHard:
		// More DFS-like behavior creates longer, winding paths (harder)
		return SelectionStrategy{
			NewestWeight: 0.7, // More DFS behavior
			OldestWeight: 0.1, // Less BFS behavior
			RandomWeight: 0.2, // Some randomness
		}
	case DifficultyExtreme:
		// Almost pure DFS creates very long, twisting paths with many dead ends
		return SelectionStrategy{
			NewestWeight: 0.9,  // Heavy DFS behavior
			OldestWeight: 0.05, // Minimal BFS behavior
			RandomWeight: 0.05, // Minimal randomness
		}
	default:
		// Default to medium difficulty
		return getStrategyForDifficulty(DifficultyMedium)
	}
}

// GetSupportedDifficulties returns a list of supported difficulty levels
func GetSupportedDifficulties() []string {
	return []string{string(DifficultyEasy), string(DifficultyMedium), string(DifficultyHard), string(DifficultyExtreme)}
}
