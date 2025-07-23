package maze

import (
	"fmt"
	"math/rand"
)

// Algorithm defines the interface for maze generation algorithms
type Algorithm interface {
	// Generate carves paths in the maze starting from the given position
	Generate(maze *Maze, startRow, startCol int, rng *rand.Rand)
}

// NewAlgorithm creates an algorithm instance by name
func NewAlgorithm(algorithmName string) (Algorithm, error) {
	switch algorithmName {
	case "dfs":
		return &DFSAlgorithm{}, nil
	default:
		return nil, fmt.Errorf("unknown algorithm: %s (supported: dfs)", algorithmName)
	}
}

// GetSupportedAlgorithms returns a list of supported algorithm names
func GetSupportedAlgorithms() []string {
	return []string{"dfs"}
}
