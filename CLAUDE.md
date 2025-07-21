# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a feature-complete CLI application for generating ASCII mazes written in Go. The project follows Test-Driven Development (TDD) principles and implements a proper Depth-First Search maze generation algorithm with advanced CLI features.

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
./maze                           # Default 21x21 maze
./maze --size 15                 # Custom size maze
./maze --seed 123 --size 9       # Reproducible maze with seed
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
go fmt ./...                     # Format code
go vet ./...                     # Lint code
```

## Architecture

The codebase follows a clean package structure with clear separation of concerns:

### Core Files
- **`main.go`**: Entry point with CLI argument parsing using Go's `flag` package
  - Handles `--size/-s` flag for maze dimensions (odd numbers, minimum 5)
  - Handles `--seed` flag for reproducible generation (string/integer)
  - Input validation and error handling with user-friendly messages

- **`internal/maze/generator.go`**: Core maze generation and representation
  - `Maze` struct: Contains Width, Height, Grid, StartRow/Col, GoalRow/Col
  - `Generator` struct: Implements DFS algorithm with configurable seed
  - `NewGenerator()`: Creates generator with random seed
  - `NewGeneratorWithSeed(string)`: Creates generator with specific seed
  - `Generate(width, height)`: DFS maze generation algorithm
  - `String()`: ASCII output with visual markers (● start, ○ goal)

- **`internal/maze/generator_test.go`**: Comprehensive test suite
  - `TestGenerateMaze`: Basic generation functionality
  - `TestMazeBoundaries`: Validates wall boundaries
  - `TestMazeString`: Output format validation
  - `TestMazePathConnectivity`: DFS connectivity verification
  - `TestMazeStartGoalMarkers`: Visual marker validation

- **`main_test.go`**: CLI integration testing
  - `TestSizeValidation`: Input validation testing
  - `TestMazeDimensions`: Size parameter verification
  - `TestSeedReproducibility`: Seed consistency testing
  - `TestDifferentSeedsDifferentMazes`: Seed variation testing

### Supporting Files
- **`Makefile`**: Development workflow automation
- **`TODO.md`**: Detailed development roadmap with phase tracking
- **`README.md`**: User documentation and feature overview
- **`CLAUDE.md`**: Developer guidance (this file)

### Maze Structure
```go
type Maze struct {
    Width, Height int
    Grid         [][]bool  // true = wall, false = path
    StartRow, StartCol int  // Start position (●)
    GoalRow, GoalCol   int  // Goal position (○)
}
```

## Current Implementation Status

### Completed Phases
- **Phase 1 (MVP)**: Basic maze generation ✅
- **Phase 2**: CLI features (size, seed) ✅
- **Phase 3**: DFS algorithm implementation ✅

### Key Features Implemented
- **DFS Algorithm**: Proper maze generation ensuring single path between any two points
- **Seed Support**: Reproducible mazes with string/integer seed conversion
- **Size Customization**: Configurable dimensions with validation
- **Visual Markers**: Start (●) and goal (○) positions
- **Path Connectivity**: Guaranteed connectivity validation with DFS traversal
- **Performance**: Optimized for large mazes (51x51 in ~0.01s)

## Testing Standards

The project maintains exceptional testing standards:
- **>95% test coverage** across all packages
- **TDD approach**: Tests written before implementation
- **Integration testing**: CLI functionality testing via exec.Command
- **Property-based testing**: Connectivity and validation testing
- **Performance testing**: Large maze generation benchmarking
- **Reproducibility testing**: Seed consistency validation

### Test Categories
1. **Unit Tests**: Core algorithm functionality
2. **Integration Tests**: CLI interface and flag parsing
3. **Property Tests**: Maze connectivity and validation
4. **Performance Tests**: Generation speed and memory usage

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
- Consistent code formatting with `go fmt`
- Static analysis compliance with `go vet`

## Algorithm Details

### Depth-First Search Implementation
1. Initialize grid with all walls
2. Start at position (1,1) and mark as path
3. Randomly shuffle direction array for each recursion
4. For each direction, check if target cell is valid and unvisited
5. If valid, carve path between current and target cell
6. Recursively continue from target cell
7. Backtrack when no valid neighbors exist

### Seed Handling
- String seeds converted to int64 via `strconv.ParseInt`
- Failed parsing falls back to custom string hashing function
- Seed applied to `rand.New(rand.NewSource(seed))` for reproducibility

## Dependencies

This project maintains zero external dependencies:
- Uses only Go standard library (`math/rand`, `strings`, `flag`, `fmt`, `os`, `strconv`)
- Testing uses standard `testing` package with `os/exec` for CLI testing
- No external frameworks or libraries required