package maze

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Maze struct {
	Width    int
	Height   int
	Grid     [][]bool // true = wall, false = path
	StartRow int
	StartCol int
	GoalRow  int
	GoalCol  int
}

type Generator struct {
	rand      *rand.Rand
	algorithm Algorithm
}

func NewGenerator() *Generator {
	algorithm, _ := NewAlgorithm("dfs") // Default to DFS algorithm
	return &Generator{
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())), // #nosec G404 - not for cryptographic use
		algorithm: algorithm,
	}
}

// NewGeneratorWithAlgorithm creates a Generator with a specific algorithm
func NewGeneratorWithAlgorithm(algorithmName string) (*Generator, error) {
	algorithm, err := NewAlgorithm(algorithmName)
	if err != nil {
		return nil, err
	}
	return &Generator{
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())), // #nosec G404 - not for cryptographic use
		algorithm: algorithm,
	}, nil
}

// NewGeneratorWithSeed creates a Generator with a specific seed for reproducible generation
func NewGeneratorWithSeed(seedStr string) *Generator {
	// Convert string seed to int64
	seed, err := strconv.ParseInt(seedStr, 10, 64)
	if err != nil {
		// If parsing fails, use string hash as fallback
		seed = hashString(seedStr)
	}

	algorithm, _ := NewAlgorithm("dfs") // Default to DFS algorithm
	return &Generator{
		rand:      rand.New(rand.NewSource(seed)), // #nosec G404 - not for cryptographic use
		algorithm: algorithm,
	}
}

// NewGeneratorWithSeedAndAlgorithm creates a Generator with specific seed and algorithm
func NewGeneratorWithSeedAndAlgorithm(seedStr, algorithmName string) (*Generator, error) {
	// Convert string seed to int64
	seed, err := strconv.ParseInt(seedStr, 10, 64)
	if err != nil {
		// If parsing fails, use string hash as fallback
		seed = hashString(seedStr)
	}

	algorithm, err := NewAlgorithm(algorithmName)
	if err != nil {
		return nil, err
	}

	return &Generator{
		rand:      rand.New(rand.NewSource(seed)), // #nosec G404 - not for cryptographic use
		algorithm: algorithm,
	}, nil
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

	// Use selected algorithm to generate maze
	g.algorithm.Generate(maze, 1, 1, g.rand)

	// Ensure start and goal positions are paths
	maze.Grid[maze.StartRow][maze.StartCol] = false
	maze.Grid[maze.GoalRow][maze.GoalCol] = false

	return maze
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
