# go-maze

A feature-rich command-line maze generator written in Go with multiple generation algorithms.

## Version 2.0.0

Complete maze generation with 5 algorithms (DFS, Kruskal's, Wilson's, Hunt-Kill, Growing Tree), multiple output formats (ASCII, Unicode, JSON), customizable difficulty levels, maze complexity analysis, customizable size, reproducible seeds, visual markers, and solution path display.

### Features

- **5 Generation Algorithms**: DFS, Kruskal's, Wilson's, Hunt-Kill, and Growing Tree algorithms
- **Algorithm selection** with `-a, --algorithm` flag (dfs, kruskal, wilson, hunt-kill, growing-tree)
- **Difficulty levels** with `-d, --difficulty` flag for Growing Tree algorithm (easy, medium, hard, extreme)
- **Maze complexity analysis** with `--analyze` flag providing detailed complexity metrics
- **Multiple output formats**: ASCII, Unicode box-drawing, and JSON with `-f, --format` flag
- **JSON analysis integration**: Analysis data included in JSON output when both flags are used
- **Solution path display** with `--solution` flag using BFS pathfinding
- **Customizable size** with `-s, --size` flag (odd numbers, minimum 5)
- **Reproducible mazes** with `--seed` flag for consistent output across all algorithms
- **Visual markers**: Start (●/◉), Goal (○/◎), and Solution path (·/•) markers
- **Path connectivity**: Guaranteed single path between any two points
- **Fast performance**: Generates 51x51 mazes in ~0.01s for all algorithms
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
./maze -a hunt-kill --size 15  # Hunt-Kill algorithm
./maze -a growing-tree --size 15  # Growing Tree algorithm

# Use different difficulty levels (Growing Tree only)
./maze -a growing-tree -d easy --size 11     # Easy difficulty
./maze -a growing-tree -d medium --size 11   # Medium difficulty (default)
./maze -a growing-tree -d hard --size 11     # Hard difficulty
./maze -a growing-tree -d extreme --size 11  # Extreme difficulty

# Use different output formats
./maze -f ascii --size 11      # ASCII format (default)
./maze -f unicode --size 11    # Unicode box-drawing characters
./maze -f json --size 11       # JSON format for programmatic use

# Display maze complexity analysis
./maze --analyze --size 11 --seed 123
./maze -a growing-tree -d hard --analyze --size 11
./maze --analyze --format json --size 11  # Analysis in JSON format

# Display solution path
./maze --solution --size 11 --seed 123
./maze -a kruskal --solution --seed 42 --size 9
./maze -f unicode --solution --seed 42 --size 9

# Combined features
./maze -a hunt-kill --solution --analyze --seed 42 --size 11
./maze -a growing-tree -d extreme --analyze --format json --size 11
./maze -a wilson --solution --analyze --format unicode --size 11

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

**Hunt-Kill Algorithm:**
```bash
./maze -a hunt-kill --seed 789 -s 11
```
```
###########
#●      # #
# ##### # #
#     # # #
##### # # #
#   # #   #
# # # ### #
# #   #   #
# ### # # #
#     #  ○#
###########
```

**Growing Tree Algorithm (Hard Difficulty):**
```bash
./maze -a growing-tree -d hard --seed 111 -s 11
```
```
###########
#●        #
# ###### ##
#      #  #
###### #  #
#    # #  #
# ## #    #
# ## #### #
#    #    #
#### #   ○#
###########
```

**Maze Analysis Output:**
```bash
./maze -a hunt-kill --analyze --seed 42 -s 9
```
```
Maze Analysis Report:

Basic Metrics:
  Total Cells: 81
  Path Cells: 31 (38.3%)
  Wall Cells: 50 (61.7%)

Complexity Metrics:
  Dead Ends: 2 (6.5% of paths)
  Junctions: 0 (0.0% of paths)
  Corridors: 29 (93.5% of paths)

Solution Metrics:
  Solution Length: 13 steps
  Optimal Ratio: 41.9% (solution vs all paths)

Difficulty Assessment:
  Difficulty Score: 45.8/100
  Difficulty Level: Medium

#########
#●  #   #
# ### # #
# #   # #
# # #####
# #     #
# ##### #
#      ○#
#########
```

**JSON Format with Analysis:**
```bash
./maze -a growing-tree -d extreme --analyze --format json --seed 42 -s 9
```
```json
{
  "width": 9,
  "height": 9,
  "grid": [...],
  "start": {"Row": 1, "Col": 1},
  "goal": {"Row": 7, "Col": 7},
  "solution_path": [],
  "analysis": {
    "total_cells": 81,
    "path_cells": 31,
    "wall_cells": 50,
    "path_density": 0.38271604938271603,
    "dead_ends": 3,
    "junctions": 1,
    "corridors": 27,
    "dead_end_ratio": 0.0967741935483871,
    "junction_ratio": 0.03225806451612903,
    "solution_length": 17,
    "optimal_ratio": 0.5483870967741935,
    "difficulty_score": 53.47,
    "difficulty_level": "Medium"
  }
}
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
  - `hunt_kill.go`: Hunt-Kill algorithm implementation for high-difficulty mazes
  - `growing_tree.go`: Growing Tree algorithm with configurable difficulty levels
  - `analyzer.go`: Maze complexity analysis system with scoring algorithm
  - `pathfinder.go`: BFS pathfinding for solution display
  - `renderer.go`: Renderer interface and factory pattern
  - `ascii_renderer.go`: ASCII format renderer (default)
  - `unicode_renderer.go`: Unicode box-drawing renderer
  - `json_renderer.go`: JSON format renderer with analysis integration
  - `*_test.go`: Comprehensive test suites with connectivity, reproducibility, snapshot, and integration testing
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

**Hunt-Kill Algorithm Maze Generation:**
1. Start with grid of all walls
2. Begin at starting position (kill phase)
3. Kill phase: Random walk carving paths until hitting dead end
4. Hunt phase: Scan grid for unvisited cells adjacent to maze
5. When found, connect to maze and start new kill phase
6. Repeat hunt-kill cycles until all cells are visited
7. Creates mazes with long winding passages and complex structure

**Growing Tree Algorithm Maze Generation:**
1. Start with grid of all walls and active cell list
2. Select cells from active list using configurable strategy weights:
   - Newest cell (DFS-like behavior) - creates long passages
   - Oldest cell (BFS-like behavior) - creates wide, branching mazes
   - Random cell - adds unpredictability
3. Difficulty levels adjust strategy weights:
   - Easy: More BFS-like (60% oldest, 20% newest, 20% random)
   - Medium: Balanced (30% oldest, 40% newest, 30% random)
   - Hard: More DFS-like (10% oldest, 70% newest, 20% random)
   - Extreme: Almost pure DFS (5% oldest, 90% newest, 5% random)
4. Continue until active list is empty

**All algorithms ensure:**
- **Perfect maze**: Exactly one path between any two points
- **No isolated areas**: All path cells are connected
- **Randomness**: Different seed values produce different mazes
- **Reproducibility**: Same seed always produces identical maze
- **Different patterns**: Each algorithm creates distinct maze characteristics

### CLI Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--algorithm` | `-a` | dfs | Algorithm for maze generation (dfs, kruskal, wilson, hunt-kill, growing-tree) |
| `--difficulty` | `-d` | medium | Difficulty level for growing-tree algorithm (easy, medium, hard, extreme) |
| `--format` | `-f` | ascii | Output format (ascii, unicode, json) |
| `--size` | `-s` | 21 | Size of square maze (must be odd, minimum 5) |
| `--seed` | - | random | Seed for reproducible generation (string/integer) |
| `--solution` | - | false | Display the solution path from start to goal |
| `--analyze` | - | false | Display detailed maze analysis and complexity metrics |
| `--help` | `-h` | - | Show help message |

### Completed Features

- [x] **5 Algorithm implementations**: DFS, Kruskal's, Wilson's, Hunt-Kill, and Growing Tree algorithms
- [x] **Algorithm selection**: CLI flag support for all 5 algorithms
- [x] **Difficulty configuration**: Growing Tree algorithm with 4 difficulty levels (easy/medium/hard/extreme)
- [x] **Maze complexity analysis**: Comprehensive analysis system with detailed metrics
- [x] **Analysis integration**: Seamless analysis display for all algorithms and formats
- [x] **Multiple output formats**: ASCII, Unicode box-drawing, and JSON formats
- [x] **Format selection**: CLI flag support with validation for output format choice
- [x] **JSON analysis integration**: Analysis data included in JSON output when requested
- [x] **Solution path display**: BFS pathfinding with `--solution` flag
- [x] **Size specification**: Custom maze dimensions with validation
- [x] **Seed support**: Reproducible maze generation for all algorithms and difficulty levels
- [x] **Visual markers**: Start (●/◉), goal (○/◎), and solution path (·/•) positions
- [x] **Path connectivity**: Guaranteed single path with comprehensive validation
- [x] **Performance optimization**: Fast generation for large mazes across all algorithms
- [x] **Comprehensive testing**: >95% test coverage with TDD approach, integration tests, and snapshot testing
- [x] **Union-Find structure**: Efficient cycle detection for Kruskal's algorithm
- [x] **Unicode rendering**: Connection-aware box-drawing character selection
- [x] **JSON output**: Structured data export with optional analysis integration

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
