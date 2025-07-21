# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a CLI application for generating ASCII mazes written in Go. The project follows Test-Driven Development (TDD) principles and is currently in MVP phase with a working maze generator.

## Commands

### Build and Run
```bash
go build -o maze .
./maze                    # Generates a 21x21 maze
```

### Testing
```bash
go test ./...             # Run all tests
go test ./internal/maze   # Run maze package tests only
go test -v ./internal/maze # Run with verbose output
```

### Development
```bash
go mod tidy              # Clean up dependencies
go fmt ./...             # Format code
```

## Architecture

The codebase follows a clean package structure:

- `main.go` - Entry point, handles CLI interface
- `internal/maze/` - Core maze generation logic
  - `generator.go` - Contains `Maze` struct and `Generator` with `Generate()` method
  - `generator_test.go` - Comprehensive test suite with 3 test cases

The `Maze` struct represents the maze state while `Generator` handles the generation algorithm. Currently generates 21x21 mazes with ASCII output using '#' for walls and ' ' for paths.

## Development Phases

The project follows a structured development approach defined in TODO.md:

- **Phase 1 (MVP)**: Basic maze generation - COMPLETED
- **Phase 2**: CLI features (custom dimensions, output formats)
- **Phase 3**: Advanced algorithms and optimizations

## Testing Standards

The project maintains high testing standards:
- All core functionality must have test coverage
- Tests validate both successful generation and edge cases
- Uses standard Go testing framework (no external test dependencies)

## Dependencies

This project uses only Go standard library - no external dependencies required.