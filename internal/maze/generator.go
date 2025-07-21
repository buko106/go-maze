package maze

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Maze struct {
	Width  int
	Height int
	Grid   [][]bool // true = wall, false = path
	StartRow int
	StartCol int
	GoalRow  int
	GoalCol  int
}

type Generator struct {
	rand *rand.Rand
}

func NewGenerator() *Generator {
	return &Generator{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// NewGeneratorWithSeed creates a Generator with a specific seed for reproducible generation
func NewGeneratorWithSeed(seedStr string) *Generator {
	// Convert string seed to int64
	seed, err := strconv.ParseInt(seedStr, 10, 64)
	if err != nil {
		// If parsing fails, use string hash as fallback
		seed = hashString(seedStr)
	}
	
	return &Generator{
		rand: rand.New(rand.NewSource(seed)),
	}
}

// hashString converts string to int64 for seed
func hashString(s string) int64 {
	var hash int64
	for _, char := range s {
		hash = hash*31 + int64(char)
	}
	return hash
}

func (g *Generator) Generate(width, height int) *Maze {
	// Initialize grid with all walls
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
		for j := range grid[i] {
			grid[i][j] = true // Start with all walls
		}
	}
	
	maze := &Maze{
		Width:    width,
		Height:   height,
		Grid:     grid,
		StartRow: 1,
		StartCol: 1,
		GoalRow:  height - 2,
		GoalCol:  width - 2,
	}
	
	// Use DFS algorithm to generate maze
	g.generateDFS(maze, 1, 1)
	
	// Ensure start and goal positions are paths
	maze.Grid[maze.StartRow][maze.StartCol] = false
	maze.Grid[maze.GoalRow][maze.GoalCol] = false
	
	return maze
}

// generateDFS implements Depth-First Search maze generation algorithm
func (g *Generator) generateDFS(maze *Maze, startRow, startCol int) {
	// Mark starting cell as path
	maze.Grid[startRow][startCol] = false
	
	// Define directions: up, right, down, left
	directions := [][2]int{{-2, 0}, {0, 2}, {2, 0}, {0, -2}}
	
	// Shuffle directions for randomness
	g.shuffleDirections(directions)
	
	// Try each direction
	for _, dir := range directions {
		newRow := startRow + dir[0]
		newCol := startCol + dir[1]
		
		// Check if new position is valid and unvisited
		if g.isValidCell(maze, newRow, newCol) && maze.Grid[newRow][newCol] {
			// Remove wall between current and new cell
			wallRow := startRow + dir[0]/2
			wallCol := startCol + dir[1]/2
			maze.Grid[wallRow][wallCol] = false
			
			// Recursively generate from new cell
			g.generateDFS(maze, newRow, newCol)
		}
	}
}

// shuffleDirections randomizes the order of directions
func (g *Generator) shuffleDirections(directions [][2]int) {
	for i := len(directions) - 1; i > 0; i-- {
		j := g.rand.Intn(i + 1)
		directions[i], directions[j] = directions[j], directions[i]
	}
}

// isValidCell checks if a cell position is within bounds
func (g *Generator) isValidCell(maze *Maze, row, col int) bool {
	return row > 0 && row < maze.Height-1 && col > 0 && col < maze.Width-1
}

func (m *Maze) String() string {
	var sb strings.Builder
	for i, row := range m.Grid {
		for j, cell := range row {
			if i == m.StartRow && j == m.StartCol {
				sb.WriteRune('●') // Filled circle for start
			} else if i == m.GoalRow && j == m.GoalCol {
				sb.WriteRune('○') // Empty circle for goal
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
