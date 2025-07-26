package maze

import (
	"fmt"
)

// MazeAnalysis contains comprehensive metrics about a maze's complexity and characteristics
type MazeAnalysis struct {
	// Basic metrics
	TotalCells  int     `json:"total_cells"`  // Total number of cells in the maze
	PathCells   int     `json:"path_cells"`   // Number of path cells (not walls)
	WallCells   int     `json:"wall_cells"`   // Number of wall cells
	PathDensity float64 `json:"path_density"` // Ratio of path cells to total cells

	// Complexity metrics
	DeadEnds      int     `json:"dead_ends"`      // Number of dead-end cells (paths with only one neighbor)
	Junctions     int     `json:"junctions"`      // Number of junction cells (paths with 3+ neighbors)
	Corridors     int     `json:"corridors"`      // Number of corridor cells (paths with exactly 2 neighbors)
	DeadEndRatio  float64 `json:"dead_end_ratio"` // Ratio of dead ends to path cells
	JunctionRatio float64 `json:"junction_ratio"` // Ratio of junctions to path cells

	// Solution metrics
	SolutionLength int     `json:"solution_length"` // Length of shortest solution path
	OptimalRatio   float64 `json:"optimal_ratio"`   // Ratio of solution length to total path cells

	// Difficulty assessment
	DifficultyScore float64 `json:"difficulty_score"` // Overall difficulty score (0-100)
	DifficultyLevel string  `json:"difficulty_level"` // Human-readable difficulty level
}

// MazeAnalyzer provides methods to analyze maze characteristics and complexity
type MazeAnalyzer struct{}

// NewAnalyzer creates a new maze analyzer instance
func NewAnalyzer() *MazeAnalyzer {
	return &MazeAnalyzer{}
}

// Analyze performs comprehensive analysis of a maze and returns detailed metrics
func (a *MazeAnalyzer) Analyze(maze *Maze) (*MazeAnalysis, error) {
	if maze == nil {
		return nil, fmt.Errorf("maze cannot be nil")
	}

	analysis := &MazeAnalysis{}

	// Calculate basic metrics
	a.calculateBasicMetrics(maze, analysis)

	// Calculate complexity metrics
	a.calculateComplexityMetrics(maze, analysis)

	// Calculate solution metrics
	a.calculateSolutionMetrics(maze, analysis)

	// Calculate overall difficulty
	a.calculateDifficultyScore(analysis)

	return analysis, nil
}

// calculateBasicMetrics computes basic maze statistics
func (a *MazeAnalyzer) calculateBasicMetrics(maze *Maze, analysis *MazeAnalysis) {
	analysis.TotalCells = maze.Width * maze.Height
	analysis.PathCells = 0
	analysis.WallCells = 0

	for i := 0; i < maze.Height; i++ {
		for j := 0; j < maze.Width; j++ {
			if maze.Grid[i][j] {
				analysis.WallCells++
			} else {
				analysis.PathCells++
			}
		}
	}

	if analysis.TotalCells > 0 {
		analysis.PathDensity = float64(analysis.PathCells) / float64(analysis.TotalCells)
	}
}

// calculateComplexityMetrics computes maze complexity statistics
func (a *MazeAnalyzer) calculateComplexityMetrics(maze *Maze, analysis *MazeAnalysis) {
	analysis.DeadEnds = 0
	analysis.Junctions = 0
	analysis.Corridors = 0

	// Analyze each path cell to determine its type
	for i := 0; i < maze.Height; i++ {
		for j := 0; j < maze.Width; j++ {
			if !maze.Grid[i][j] { // It's a path cell
				neighborCount := a.countPathNeighbors(maze, i, j)

				switch neighborCount {
				case 1:
					analysis.DeadEnds++
				case 2:
					analysis.Corridors++
				case 3, 4:
					analysis.Junctions++
				}
			}
		}
	}

	// Calculate ratios
	if analysis.PathCells > 0 {
		analysis.DeadEndRatio = float64(analysis.DeadEnds) / float64(analysis.PathCells)
		analysis.JunctionRatio = float64(analysis.Junctions) / float64(analysis.PathCells)
	}
}

// calculateSolutionMetrics computes solution-related statistics
func (a *MazeAnalyzer) calculateSolutionMetrics(maze *Maze, analysis *MazeAnalysis) {
	// Use pathfinder to get solution path
	solutionPath := FindPath(maze)

	analysis.SolutionLength = len(solutionPath)

	if analysis.PathCells > 0 && analysis.SolutionLength > 0 {
		analysis.OptimalRatio = float64(analysis.SolutionLength) / float64(analysis.PathCells)
	}
}

// calculateDifficultyScore computes an overall difficulty score and level
func (a *MazeAnalyzer) calculateDifficultyScore(analysis *MazeAnalysis) {
	// Difficulty is based on multiple factors:
	// - High dead-end ratio increases difficulty (more wrong turns)
	// - Low path density increases difficulty (fewer options)
	// - High solution length ratio increases difficulty (longer optimal path)
	// - Low junction ratio can increase difficulty (fewer decision points, but longer corridors)

	score := 0.0

	// Dead-end factor (0-30 points): More dead ends = higher difficulty
	deadEndFactor := analysis.DeadEndRatio * 30.0
	if deadEndFactor > 30.0 {
		deadEndFactor = 30.0
	}
	score += deadEndFactor

	// Path density factor (0-20 points): Lower density = higher difficulty
	densityFactor := (1.0 - analysis.PathDensity) * 20.0
	if densityFactor > 20.0 {
		densityFactor = 20.0
	}
	score += densityFactor

	// Solution length factor (0-25 points): Longer solution = higher difficulty
	if analysis.OptimalRatio > 0 {
		solutionFactor := analysis.OptimalRatio * 25.0
		if solutionFactor > 25.0 {
			solutionFactor = 25.0
		}
		score += solutionFactor
	}

	// Junction factor (0-15 points): Fewer junctions can mean longer corridors
	junctionFactor := (1.0 - analysis.JunctionRatio) * 15.0
	if junctionFactor > 15.0 {
		junctionFactor = 15.0
	}
	score += junctionFactor

	// Complexity factor (0-10 points): Balance of different cell types
	if analysis.PathCells > 0 {
		varietyScore := 0.0
		if analysis.DeadEnds > 0 {
			varietyScore += 3.0
		}
		if analysis.Junctions > 0 {
			varietyScore += 4.0
		}
		if analysis.Corridors > 0 {
			varietyScore += 3.0
		}
		score += varietyScore
	}

	analysis.DifficultyScore = score

	// Assign difficulty level based on score
	switch {
	case score >= 75.0:
		analysis.DifficultyLevel = "Extreme"
	case score >= 55.0:
		analysis.DifficultyLevel = "Hard"
	case score >= 35.0:
		analysis.DifficultyLevel = "Medium"
	case score >= 15.0:
		analysis.DifficultyLevel = "Easy"
	default:
		analysis.DifficultyLevel = "Trivial"
	}
}

// countPathNeighbors counts the number of adjacent path cells (not diagonal)
func (a *MazeAnalyzer) countPathNeighbors(maze *Maze, row, col int) int {
	count := 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		// Check bounds and if it's a path
		if newRow >= 0 && newRow < maze.Height &&
			newCol >= 0 && newCol < maze.Width &&
			!maze.Grid[newRow][newCol] {
			count++
		}
	}

	return count
}

// String returns a formatted string representation of the analysis
func (a *MazeAnalysis) String() string {
	return fmt.Sprintf(`Maze Analysis Report:

Basic Metrics:
  Total Cells: %d
  Path Cells: %d (%.1f%%)
  Wall Cells: %d (%.1f%%)

Complexity Metrics:
  Dead Ends: %d (%.1f%% of paths)
  Junctions: %d (%.1f%% of paths)
  Corridors: %d (%.1f%% of paths)

Solution Metrics:
  Solution Length: %d steps
  Optimal Ratio: %.1f%% (solution vs all paths)

Difficulty Assessment:
  Difficulty Score: %.1f/100
  Difficulty Level: %s`,
		a.TotalCells,
		a.PathCells, a.PathDensity*100,
		a.WallCells, (1-a.PathDensity)*100,
		a.DeadEnds, a.DeadEndRatio*100,
		a.Junctions, a.JunctionRatio*100,
		a.Corridors, (1-a.DeadEndRatio-a.JunctionRatio)*100,
		a.SolutionLength,
		a.OptimalRatio*100,
		a.DifficultyScore,
		a.DifficultyLevel)
}
