package maze

import "math/rand"

// KruskalAlgorithm implements maze generation using Kruskal's algorithm
type KruskalAlgorithm struct{}

// Edge represents a connection between two cells in the maze
type Edge struct {
	fromRow, fromCol int
	toRow, toCol     int
}

// UnionFind data structure for tracking connected components
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind creates a new union-find structure
func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

// Find returns the root of the component containing x
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

// Union merges the components containing x and y
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false // Already in same component
	}

	// Union by rank
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}

// Generate implements the Algorithm interface using Kruskal's algorithm
func (k *KruskalAlgorithm) Generate(maze *Maze, startRow, startCol int, rng *rand.Rand) {
	// Create all possible edges between adjacent cells
	edges := k.createEdges(maze)

	// Shuffle edges for randomness
	k.shuffleEdges(edges, rng)

	// Create union-find structure for cells
	cellCount := ((maze.Height - 1) / 2) * ((maze.Width - 1) / 2)
	uf := NewUnionFind(cellCount)

	// Process edges and connect components
	for _, edge := range edges {
		cell1 := k.cellToIndex(maze, edge.fromRow, edge.fromCol)
		cell2 := k.cellToIndex(maze, edge.toRow, edge.toCol)

		// If cells are in different components, connect them
		if uf.Union(cell1, cell2) {
			// Mark both cells as paths
			maze.Grid[edge.fromRow][edge.fromCol] = false
			maze.Grid[edge.toRow][edge.toCol] = false

			// Remove wall between cells
			wallRow := (edge.fromRow + edge.toRow) / 2
			wallCol := (edge.fromCol + edge.toCol) / 2
			maze.Grid[wallRow][wallCol] = false
		}
	}
}

// createEdges generates all possible edges between adjacent cells
func (k *KruskalAlgorithm) createEdges(maze *Maze) []Edge {
	var edges []Edge

	// Create edges between horizontally adjacent cells
	for row := 1; row < maze.Height-1; row += 2 {
		for col := 1; col < maze.Width-3; col += 2 {
			edges = append(edges, Edge{
				fromRow: row,
				fromCol: col,
				toRow:   row,
				toCol:   col + 2,
			})
		}
	}

	// Create edges between vertically adjacent cells
	for row := 1; row < maze.Height-3; row += 2 {
		for col := 1; col < maze.Width-1; col += 2 {
			edges = append(edges, Edge{
				fromRow: row,
				fromCol: col,
				toRow:   row + 2,
				toCol:   col,
			})
		}
	}

	return edges
}

// shuffleEdges randomizes the order of edges
func (k *KruskalAlgorithm) shuffleEdges(edges []Edge, rng *rand.Rand) {
	for i := len(edges) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		edges[i], edges[j] = edges[j], edges[i]
	}
}

// cellToIndex converts cell coordinates to a unique index
func (k *KruskalAlgorithm) cellToIndex(maze *Maze, row, col int) int {
	cellRow := (row - 1) / 2
	cellCol := (col - 1) / 2
	cellsPerRow := (maze.Width - 1) / 2
	return cellRow*cellsPerRow + cellCol
}
