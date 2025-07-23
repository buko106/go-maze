package maze

// Position represents a coordinate in the maze
type Position struct {
	Row int
	Col int
}

// FindPath finds the shortest path from start to goal using BFS
func FindPath(maze *Maze) []Position {
	if maze.Grid[maze.StartRow][maze.StartCol] || maze.Grid[maze.GoalRow][maze.GoalCol] {
		return nil // Start or goal is blocked
	}

	// BFS to find shortest path
	queue := []Position{{Row: maze.StartRow, Col: maze.StartCol}}
	visited := make(map[Position]bool)
	parent := make(map[Position]Position)

	visited[Position{Row: maze.StartRow, Col: maze.StartCol}] = true

	// Directions: up, right, down, left
	directions := []Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check if we reached the goal
		if current.Row == maze.GoalRow && current.Col == maze.GoalCol {
			// Reconstruct path
			path := []Position{}
			for pos := current; pos != (Position{Row: maze.StartRow, Col: maze.StartCol}); pos = parent[pos] {
				path = append([]Position{pos}, path...)
			}
			path = append([]Position{{Row: maze.StartRow, Col: maze.StartCol}}, path...)
			return path
		}

		// Explore neighbors
		for _, dir := range directions {
			newPos := Position{
				Row: current.Row + dir.Row,
				Col: current.Col + dir.Col,
			}

			// Check bounds and if it's a valid path
			if newPos.Row >= 0 && newPos.Row < maze.Height &&
				newPos.Col >= 0 && newPos.Col < maze.Width &&
				!maze.Grid[newPos.Row][newPos.Col] && // Not a wall
				!visited[newPos] {

				visited[newPos] = true
				parent[newPos] = current
				queue = append(queue, newPos)
			}
		}
	}

	return nil // No path found
}
