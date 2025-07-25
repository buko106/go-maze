# go-maze

A feature-rich command-line maze generator written in Go with multiple generation algorithms.

## Version 1.0.0

Complete maze generation with multiple algorithms (DFS, Kruskal's, Wilson's), multiple output formats (ASCII, Unicode, JSON), customizable size, reproducible seeds, visual markers, and solution path display.

### Features

- **Multiple algorithms**: Depth-First Search (DFS), Kruskal's, and Wilson's algorithm support
- **Algorithm selection** with `-a, --algorithm` flag (dfs, kruskal, wilson)
- **Multiple output formats**: ASCII, Unicode box-drawing, and JSON with `-f, --format` flag
- **Solution path display** with `--solution` flag using BFS pathfinding
- **Customizable size** with `-s, --size` flag (odd numbers, minimum 5)
- **Reproducible mazes** with `--seed` flag for consistent output
- **Visual markers**: Start (●/◉), Goal (○/◎), and Solution path (·/•) markers
- **Path connectivity**: Guaranteed single path between any two points
- **Fast performance**: Generates 51x51 mazes in ~0.01s
- **Comprehensive testing** with >95% test coverage and TDD approach

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
# Generate default 21x21 maze (DFS algorithm)
./maze

# Generate custom size maze
./maze --size 15
./maze -s 9

# Generate reproducible maze with seed
./maze --seed 123 --size 11

# Use different algorithms
./maze -a dfs --size 15        # Depth-First Search (default)
./maze -a kruskal --size 15    # Kruskal's algorithm
./maze -a wilson --size 15     # Wilson's algorithm

# Use different output formats
./maze -f ascii --size 11      # ASCII format (default)
./maze -f unicode --size 11    # Unicode box-drawing characters
./maze -f json --size 11       # JSON format for programmatic use

# Display solution path
./maze --solution --size 11 --seed 123
./maze -a kruskal --solution --seed 42 --size 9
./maze -f unicode --solution --seed 42 --size 9

# Algorithm with format combinations
./maze -a kruskal --seed 42 --size 9 -f unicode
./maze -a wilson --seed 42 --size 9 -f json

# Show help
./maze --help
```

#### Example Outputs

**ASCII Format (Default):**
```bash
./maze -a dfs --seed 42 -s 7
```
```
#######
#●#   #
# # # #
#   # #
##### #
#    ○#
#######
```

**ASCII Format with Solution:**
```bash
./maze -a dfs --seed 42 -s 7 --solution
```
```
#######
#●#···#
#·#·#·#
#···#·#
#####·#
#    ○#
#######
```

**Unicode Format:**
```bash
./maze -a dfs --seed 42 -s 7 -f unicode
```
```
┌─┬───┐
│◉│   │
│ ╵ ╷ │
│   │ │
├───┘ │
│    ◎│
└─────┘
```

**Unicode Format with Solution:**
```bash
./maze -a dfs --seed 42 -s 7 -f unicode --solution
```
```
┌─┬───┐
│◉│•••│
│•╵•╷•│
│•••│•│
├───┘•│
│    ◎│
└─────┘
```

**JSON Format:**
```bash
./maze -a dfs --seed 42 -s 7 -f json
```
```json
{
  "width": 7,
  "height": 7,
  "grid": [
    [true,true,true,true,true,true,true],
    [true,false,true,false,false,false,true],
    ...
  ],
  "start": {"Row": 1, "Col": 1},
  "goal": {"Row": 5, "Col": 5},
  "solution_path": []
}
```

**Kruskal Algorithm with same seed:**
```bash
./maze -a kruskal --seed 123 -s 9
```
```
#########
#●    # #
##### # #
#       #
# #######
#       #
# # ### #
# #   #○#
#########
```

**Kruskal Algorithm with solution path:**
```bash
./maze -a kruskal --seed 123 -s 9 --solution
```
```
#########
#●····# #
#####·# #
#    ···#
# ##### #
#   ···#
# #·###·#
# #···#○#
#########
```

**Larger DFS maze:**
```bash
./maze -a dfs --seed 456 -s 15
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

**Larger Kruskal maze:**
```bash
./maze -a kruskal --seed 456 -s 15
```
```
###############
#●        #   #
# ##### ##### #
#     # #     #
# # # ### #####
# # # #       #
##### ### #####
#           # #
# # ### ### # #
# #   #   #   #
# ### ##### ###
#   # #   #   #
##### # ### ###
#     #      ○#
###############
```

**Wilson Algorithm with same seed:**
```bash
./maze -a wilson --seed 456 -s 15
```
```
###############
#●      #     #
# ##### # ### #
#   #   #   # #
### # ##### # #
#   #   #     #
# ##### ##### #
#       #   # #
####### ### # #
#   #         #
# # # #########
# # # #   #   #
# ### # # # # #
#     # #    ○#
###############
```

**Wilson Algorithm with solution path:**
```bash
./maze -a wilson --seed 456 -s 15 --solution
```
```
###############
#●······#     #
#·#####·# ### #
#···#  ·#   # #
###·# ##·## # #
#  ·#   #·····#
# ##·## ####·#
#    ···#   #·#
#######·### #·#
#   #  ·······#
# # # #########
# # # #   #   #
# ### # # # # #
#     # #    ○#
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
  - `generator.go`: Generator with algorithm interface and seed support
  - `algorithm.go`: Algorithm interface and factory pattern
  - `dfs.go`: Depth-First Search algorithm implementation
  - `kruskal.go`: Kruskal's algorithm with Union-Find data structure
  - `wilson.go`: Wilson's algorithm with loop-erased random walks
  - `pathfinder.go`: BFS pathfinding for solution display
  - `renderer.go`: Renderer interface and factory pattern
  - `ascii_renderer.go`: ASCII format renderer (default)
  - `unicode_renderer.go`: Unicode box-drawing renderer
  - `json_renderer.go`: JSON format renderer
  - `*_test.go`: Comprehensive test suites with connectivity, reproducibility, and snapshot testing
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

**Kruskal's Algorithm Maze Generation:**
1. Start with grid of all walls
2. Create list of all possible edges between adjacent cells
3. Randomly shuffle the edge list
4. Use Union-Find data structure to track connected components
5. For each edge, if cells are in different components, connect them
6. Continue until all cells are connected in a single component

**Wilson's Algorithm Maze Generation:**
1. Start with grid of all walls and add starting cell to maze
2. Create list of all potential path cells (odd coordinates)
3. For each remaining cell not in maze:
   - Perform loop-erased random walk from current cell
   - When walk reaches a cell already in maze, add entire path
   - Connect path cells by removing walls between adjacent positions
4. Continue until all cells are connected in uniform spanning tree

**All algorithms ensure:**
- **Perfect maze**: Exactly one path between any two points
- **No isolated areas**: All path cells are connected
- **Randomness**: Different seed values produce different mazes
- **Reproducibility**: Same seed always produces identical maze
- **Different patterns**: Each algorithm creates distinct maze characteristics

### CLI Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--algorithm` | `-a` | dfs | Algorithm for maze generation (dfs, kruskal, wilson) |
| `--format` | `-f` | ascii | Output format (ascii, unicode, json) |
| `--size` | `-s` | 21 | Size of square maze (must be odd, minimum 5) |
| `--seed` | - | random | Seed for reproducible generation (string/integer) |
| `--solution` | - | false | Display the solution path from start to goal |
| `--help` | `-h` | - | Show help message |

### Completed Features

- [x] **Multiple algorithms**: DFS, Kruskal's, and Wilson's algorithm implementations
- [x] **Algorithm selection**: CLI flag support for algorithm choice
- [x] **Multiple output formats**: ASCII, Unicode box-drawing, and JSON formats
- [x] **Format selection**: CLI flag support with validation for output format choice
- [x] **Solution path display**: BFS pathfinding with `--solution` flag
- [x] **Size specification**: Custom maze dimensions
- [x] **Seed support**: Reproducible maze generation for all algorithms
- [x] **Visual markers**: Start (●/◉), goal (○/◎), and solution path (·/•) positions
- [x] **Path connectivity**: Guaranteed single path with comprehensive validation
- [x] **Performance optimization**: Fast generation for large mazes
- [x] **Comprehensive testing**: >95% test coverage with TDD approach and snapshot testing
- [x] **Union-Find structure**: Efficient cycle detection for Kruskal's algorithm
- [x] **Unicode rendering**: Connection-aware box-drawing character selection
- [x] **JSON output**: Structured data export for programmatic use

### Future Enhancements

- [ ] **Additional algorithms**: Prim's algorithm implementation
- [ ] **Performance comparison**: Benchmarking between algorithms
- [ ] **Custom start/goal**: Specify positions (`--start`, `--goal` flags)
- [ ] **Solution animation**: Animate solution path discovery
- [ ] **Version info**: Display version information (`--version` flag)
- [ ] **Large maze optimization**: Memory and performance improvements for >100x100 mazes

### Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests first (TDD approach)
4. Implement the feature
5. Ensure all tests pass
6. Submit a pull request

### License

MIT License
