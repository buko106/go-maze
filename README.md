# go-maze

A feature-rich command-line maze generator written in Go using Depth-First Search algorithm.

## Version 0.2.0

Complete maze generation with customizable size, reproducible seeds, and visual start/goal markers.

### Features

- **Proper maze generation** using Depth-First Search (DFS) algorithm
- **Customizable size** with `-s, --size` flag (odd numbers, minimum 5)
- **Reproducible mazes** with `--seed` flag for consistent output
- **Visual markers**: Start (●) and Goal (○) positions
- **Path connectivity**: Guaranteed single path between any two points
- **Fast performance**: Generates 51x51 mazes in ~0.01s
- **Comprehensive testing** with >95% test coverage

### Installation

```bash
git clone https://github.com/buko106/go-maze.git
cd go-maze
make install
```

Or using Go directly:
```bash
go build -o maze .
```

### Usage

#### Basic Usage
```bash
# Generate default 21x21 maze
./maze

# Generate custom size maze
./maze --size 15
./maze -s 9

# Generate reproducible maze with seed
./maze --seed 123 --size 11

# Show help
./maze --help
```

#### Example Outputs

**Small maze with seed:**
```bash
./maze --seed 123 -s 9
```
```
#########
#●#     #
# # ### #
#   # # #
##### # #
#   #   #
# # # ###
# #    ○#
#########
```

**Larger maze:**
```bash
./maze --seed 456 -s 15
```
```
###############
#●    #       #
### # # ##### #
#   #   #   # #
# ##### # # # #
#       # #   #
######### ### #
#   #     #   #
# # # ##### # #
# #   #   # # #
# ### ### # # #
#   #     #   #
### ####### # #
#          #○#
###############
```

### Development

This project uses **Makefile** for common development tasks and follows **Test-Driven Development (TDD)** principles.

#### Quick Start
```bash
# Development workflow (clean, format, lint, test, build)
make dev

# Build the binary
make build

# Run tests
make test

# Run with verbose output
make test-verbose

# Generate coverage report
make coverage

# Format and lint code
make fmt lint

# Run examples
make examples

# Show all available targets
make help
```

#### Manual Commands
```bash
# Build and install
go build -o maze .

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Format code
go fmt ./...
```

### Architecture

- **`main.go`**: Entry point and CLI argument parsing with `flag` package
- **`internal/maze/`**: Core maze generation and representation
  - `generator.go`: DFS algorithm implementation with seed support
  - `generator_test.go`: Comprehensive test suite with connectivity testing
- **`Makefile`**: Development workflow automation
- **`TODO.md`**: Detailed development roadmap and task tracking

### Algorithm Details

**Depth-First Search (DFS) Maze Generation:**
1. Start with grid of all walls
2. Begin at starting position (1,1)
3. Randomly select unvisited neighboring cells
4. Carve path to neighbor and remove wall between
5. Recursively continue from new cell
6. Backtrack when no unvisited neighbors remain

This ensures:
- **Perfect maze**: Exactly one path between any two points
- **No isolated areas**: All path cells are connected
- **Randomness**: Different seed values produce different mazes
- **Reproducibility**: Same seed always produces identical maze

### CLI Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--size` | `-s` | 21 | Size of square maze (must be odd, minimum 5) |
| `--seed` | - | random | Seed for reproducible generation (string/integer) |
| `--help` | `-h` | - | Show help message |

### Completed Features

- [x] **Size specification**: Custom maze dimensions
- [x] **Seed support**: Reproducible maze generation
- [x] **DFS algorithm**: Proper maze generation with guaranteed connectivity
- [x] **Visual markers**: Start (●) and goal (○) positions
- [x] **Performance optimization**: Fast generation for large mazes
- [x] **Comprehensive testing**: Full test coverage with TDD approach

### Upcoming Features

- [ ] **Algorithm selection**: Kruskal, Prim algorithms (`--algorithm` flag)
- [ ] **Output formats**: Unicode box-drawing, JSON export (`--format` flag)
- [ ] **Solution display**: Show path from start to goal (`--solution` flag)
- [ ] **Custom start/goal**: Specify positions (`--start`, `--goal` flags)
- [ ] **Version info**: Display version information (`--version` flag)

### Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests first (TDD approach)
4. Implement the feature
5. Ensure all tests pass
6. Submit a pull request

### License

MIT License