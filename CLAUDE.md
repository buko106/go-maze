# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a feature-complete CLI application for generating ASCII mazes written in Go. The project follows Test-Driven Development (TDD) principles and implements multiple maze generation algorithms (DFS, Kruskal's) with advanced CLI features including solution path display.

## Commands

### Primary Development Commands (Use Makefile)
```bash
make dev                  # Full development workflow: clean, format, lint, test, build
make build               # Build the binary to ./bin/maze
make install             # Build to current directory as ./maze
make test                # Run all tests
make test-verbose        # Run tests with verbose output
make test-coverage       # Run tests with coverage
make coverage            # Generate HTML coverage report
make examples            # Show example maze outputs
make help                # Show all available targets
```

### Direct Build and Run
```bash
go build -o maze .
./maze                           # Default 21x21 maze (DFS algorithm)
./maze --size 15                 # Custom size maze
./maze --seed 123 --size 9       # Reproducible maze with seed
./maze --algorithm kruskal --size 11  # Use Kruskal's algorithm
./maze --solution --seed 42 --size 9  # Display solution path
./maze --help                    # Show usage information
```

### Testing
```bash
go test ./...                    # Run all tests
go test ./internal/maze          # Run maze package tests only
go test -v ./internal/maze       # Run with verbose output
go test -cover ./...             # Run tests with coverage report
```

### Development Tools
```bash
go mod tidy                      # Clean up dependencies
make fmt                         # Format code using golangci-lint
make lint                        # Lint code using golangci-lint
make fmt-go                      # Format code using go fmt (legacy)
make lint-go                     # Lint code using go vet (legacy)
```

### Code Quality Tools
```bash
golangci-lint run                # Run comprehensive linting
golangci-lint fmt                # Apply formatters and fix issues
golangci-lint linters            # List available linters
golangci-lint formatters         # List available formatters
```

### Pre-commit Hooks
```bash
pre-commit install               # Install pre-commit hooks to git
pre-commit run --all-files       # Run all hooks on all files
pre-commit run <hook-id>         # Run specific hook
pre-commit autoupdate            # Update hook versions
pre-commit clean                 # Clean hook environments
```

**Enabled hooks:**
- **golangci-lint**: Lint changed files only (fast)
- **golangci-lint-fmt**: Format all Go files
- **go-mod-tidy**: Keep go.mod dependencies tidy
- **go-unit-tests**: Run tests for changed packages
- **go-build**: Verify project builds successfully
- **File checks**: YAML/JSON validation, whitespace fixes, merge conflict detection

## Architecture

The codebase follows a clean package structure with clear separation of concerns:

### Core Files
- **`main.go`**: Entry point with CLI argument parsing using Go's `flag` package
  - Handles `--size/-s` flag for maze dimensions (odd numbers, minimum 5)
  - Handles `--seed` flag for reproducible generation (string/integer)
  - Handles `--algorithm/-a` flag for algorithm selection (dfs, kruskal)
  - Handles `--solution` flag for solution path display
  - Input validation and error handling with user-friendly messages

- **`internal/maze/generator.go`**: Core maze generation and representation
  - `Maze` struct: Contains Width, Height, Grid, StartRow/Col, GoalRow/Col, SolutionPath
  - `Generator` struct: Configurable generator with algorithm selection and seeding
  - `NewGenerator()`: Creates generator with random seed and default DFS algorithm
  - `NewGeneratorWithSeed(string)`: Creates generator with specific seed
  - `NewGeneratorWithAlgorithm(string)`: Creates generator with specific algorithm
  - `Generate(width, height)`: Multi-algorithm maze generation
  - `String()`: ASCII output with visual markers (● start, ○ goal, · solution path)

- **`internal/maze/algorithm.go`**: Algorithm interface and factory pattern
  - `Algorithm` interface: Common interface for all generation algorithms
  - `NewAlgorithm(string)`: Factory function for algorithm creation
  - `GetSupportedAlgorithms()`: Lists available algorithms (dfs, kruskal)

- **`internal/maze/dfs.go`**: Depth-First Search algorithm implementation
  - `DFSAlgorithm` struct: Implements recursive DFS with random direction selection
  - Ensures perfect maze properties with single path connectivity

- **`internal/maze/kruskal.go`**: Kruskal's algorithm with Union-Find implementation
  - `KruskalAlgorithm` struct: Implements minimum spanning tree approach
  - `UnionFind` data structure: Efficient cycle detection and path compression
  - Creates mazes with different structural characteristics than DFS

- **`internal/maze/pathfinder.go`**: BFS pathfinding for solution display
  - `Position` struct: Represents coordinates in the maze
  - `FindPath(maze)`: BFS algorithm for shortest path from start to goal
  - Returns slice of positions representing the optimal solution path

- **`internal/maze/generator_test.go`**: Comprehensive test suite for generation
  - `TestGenerateMaze`: Basic generation functionality
  - `TestMazeBoundaries`: Validates wall boundaries
  - `TestMazeString`: Output format validation
  - `TestMazePathConnectivity`: DFS connectivity verification
  - `TestMazeStartGoalMarkers`: Visual marker validation
  - `TestNewGeneratorWithAlgorithm`: Algorithm selection testing
  - `TestHashStringFunctionality`: Seed hashing validation

- **`internal/maze/algorithm_test.go`**: Algorithm interface testing
  - `TestNewAlgorithm`: Algorithm factory testing
  - `TestGetSupportedAlgorithms`: Algorithm enumeration testing
  - `TestDFSAlgorithmGenerate`: DFS-specific functionality
  - `TestKruskalAlgorithmGenerate`: Kruskal-specific functionality
  - Cross-algorithm reproducibility and connectivity testing

- **`internal/maze/pathfinder_test.go`**: Pathfinding algorithm testing
  - `TestFindPathBasic`: Basic pathfinding functionality
  - `TestFindPathNoPath`: No-path scenario handling
  - `TestFindPathStartBlocked`: Blocked start position testing
  - `TestFindPathGoalBlocked`: Blocked goal position testing
  - `TestFindPathSameStartGoal`: Edge case validation
  - `TestFindPathConnectivity`: Path adjacency validation

- **`main_test.go`**: CLI integration testing
  - `TestSizeValidation`: Input validation testing
  - `TestMazeDimensions`: Size parameter verification
  - `TestSeedReproducibility`: Seed consistency testing
  - `TestAlgorithmFlag`: Algorithm selection CLI testing
  - `TestSolutionFlag`: Solution display CLI testing
  - `TestSolutionPathContinuity`: Solution path validation
  - `TestSolutionWithDifferentSeeds`: Solution variation testing

### Supporting Files
- **`Makefile`**: Development workflow automation
- **`TODO.md`**: Detailed development roadmap with phase tracking
- **`README.md`**: User documentation and feature overview
- **`CLAUDE.md`**: Developer guidance (this file)

### Maze Structure
```go
type Maze struct {
    Width        int
    Height       int
    Grid         [][]bool   // true = wall, false = path
    StartRow     int        // Start position (●)
    StartCol     int
    GoalRow      int        // Goal position (○)
    GoalCol      int
    SolutionPath []Position // Optional solution path from start to goal (·)
}

type Position struct {
    Row int
    Col int
}
```

## Current Implementation Status

### Completed Phases
- **Phase 1 (MVP)**: Basic maze generation ✅
- **Phase 2**: CLI features (size, seed) ✅
- **Phase 3**: Algorithm implementation (DFS, Kruskal's) ✅
- **Phase 4**: Solution display feature ✅

### Key Features Implemented
- **Multiple Algorithms**: DFS and Kruskal's algorithm implementations with distinct characteristics
- **Algorithm Selection**: CLI flag support for choosing generation algorithm
- **Solution Path Display**: BFS pathfinding with visual solution markers
- **Seed Support**: Reproducible mazes with string/integer seed conversion for all algorithms
- **Size Customization**: Configurable dimensions with validation (odd numbers, minimum 5)
- **Visual Markers**: Start (●), goal (○), and solution path (·) positions
- **Path Connectivity**: Guaranteed connectivity validation with algorithm-specific testing
- **Performance**: Optimized for large mazes (51x51 in ~0.01s for both algorithms)
- **Union-Find Structure**: Efficient cycle detection for Kruskal's algorithm

## Testing Standards

The project maintains exceptional testing standards:
- **>95% test coverage** across all packages
- **TDD approach**: Tests written before implementation
- **Integration testing**: CLI functionality testing via exec.Command
- **Property-based testing**: Connectivity and validation testing
- **Performance testing**: Large maze generation benchmarking
- **Reproducibility testing**: Seed consistency validation

### Test Categories
1. **Unit Tests**: Core algorithm functionality (DFS, Kruskal's, BFS pathfinding)
2. **Integration Tests**: CLI interface and flag parsing (all algorithms and solution display)
3. **Property Tests**: Maze connectivity and validation for all generation algorithms
4. **Performance Tests**: Generation speed and memory usage across algorithms
5. **Pathfinding Tests**: Solution path correctness and edge case handling

## Development Workflow

### TDD Process (Strictly Followed)
1. **Red**: Write failing test first
2. **Green**: Implement minimal code to pass test
3. **Refactor**: Improve code while keeping tests green
4. **Repeat**: Continue for each feature

### Code Quality Standards
- Go standard library only (no external dependencies)
- Comprehensive error handling with user-friendly messages
- Clean separation between CLI, algorithm, and testing concerns
- **Modern linting**: Comprehensive code analysis with `golangci-lint`
- **Multi-formatter support**: Formatting with `gofmt` and `goimports`
- **Security analysis**: Security checks with `gosec`
- **Code complexity**: Cyclomatic complexity monitoring with `gocyclo`
- **Export documentation**: Proper comments for all exported functions and types

## Algorithm Details

### Depth-First Search Implementation
1. Initialize grid with all walls
2. Start at position (1,1) and mark as path
3. Randomly shuffle direction array for each recursion
4. For each direction, check if target cell is valid and unvisited
5. If valid, carve path between current and target cell
6. Recursively continue from target cell
7. Backtrack when no valid neighbors exist

### Kruskal's Algorithm Implementation
1. Initialize grid with all walls
2. Create list of all possible edges between adjacent cells
3. Randomly shuffle the edge list using configured seed
4. Initialize Union-Find data structure for cycle detection
5. For each edge, check if connecting cells would create a cycle
6. If no cycle, connect the cells by removing the wall
7. Continue until all cells are connected in a spanning tree

### BFS Pathfinding Implementation
1. Initialize BFS queue with start position
2. Track visited cells and parent relationships
3. Explore neighbors in breadth-first order
4. When goal is reached, reconstruct path using parent pointers
5. Return shortest path from start to goal as Position slice

### Seed Handling
- String seeds converted to int64 via `strconv.ParseInt`
- Failed parsing falls back to custom string hashing function
- Seed applied to `rand.New(rand.NewSource(seed))` for reproducibility
- Same seed produces identical results across all algorithms

## Dependencies

This project maintains zero runtime dependencies:
- Uses only Go standard library (`math/rand`, `strings`, `flag`, `fmt`, `os`, `strconv`)
- Testing uses standard `testing` package with `os/exec` for CLI testing
- No external frameworks or libraries required

### Development Dependencies
- **golangci-lint v2.3.0+**: Comprehensive linting and formatting toolchain
- **pre-commit v4.2.0+**: Automated code quality checks before commits
- Configuration files: `.golangci.yml` (version 2 format), `.pre-commit-config.yaml`
- Includes 12+ linters: errcheck, govet, staticcheck, gosec, revive, etc.
- Auto-formatting with gofmt and goimports
- Pre-commit hooks for continuous quality assurance
