# go-maze Development TODO List

## Phase 1: MVP Implementation ✅ COMPLETED

### Core Functionality
- [x] **Set up basic project structure**
  - [x] Create `go.mod` with module `github.com/buko106/go-maze`
  - [x] Create main package entry point
  - [x] Set up internal package structure

- [x] **Implement basic maze generation (TDD)**
  - [x] Write test for maze creation with fixed dimensions
  - [x] Write test for maze boundary validation (all edges should be walls)
  - [x] Write test for maze string representation
  - [x] Implement `Maze` struct with `Width`, `Height`, `Grid` fields
  - [x] Implement `Generator` struct with basic random generation
  - [x] Implement `String()` method for ASCII output

- [x] **Build and execution verification**
  - [x] Ensure `go build` produces working binary
  - [x] Verify maze output is properly formatted
  - [x] Test on multiple runs to ensure randomness

### Testing Requirements
- [x] All functions must have unit tests
- [x] Test coverage should be >80%
- [x] Tests should follow TDD red-green-refactor cycle

## Phase 2: CLI Features (Current Priority)

### Command-line Interface
- [x] **Size specification feature** ✅ COMPLETED
  - [x] Add CLI argument parsing (used standard library `flag` package)
  - [x] Implement `-s, --size` flag for square mazes
  - [x] Add input validation (minimum size 5, odd numbers only)
  - [x] Write tests for various size inputs
  - [x] Handle odd size requirements for proper maze generation

- [ ] **Seed specification feature**
  - [ ] Implement `--seed` flag for reproducible mazes
  - [ ] Modify Generator to accept seed parameter
  - [ ] Write tests to verify same seed produces same maze
  - [ ] Add seed validation

- [ ] **Help and version display**
  - [x] Implement `-h, --help` flag (automatic with flag package)
  - [ ] Add `--version` flag
  - [ ] Create comprehensive usage documentation
  - [ ] Add examples in help text

### Input Validation
- [x] Size range validation (minimum 5, odd numbers only) ✅ COMPLETED
- [ ] Seed value validation
- [x] Error handling for invalid inputs ✅ COMPLETED
- [x] User-friendly error messages ✅ COMPLETED

## Phase 3: Algorithm Implementation

### Maze Generation Algorithms
- [ ] **Implement proper maze generation algorithm**
  - [ ] Replace random wall placement with Depth-First Search (DFS)
  - [ ] Ensure generated maze has exactly one path between any two points
  - [ ] Write tests to verify path connectivity
  - [ ] Performance testing for large mazes

- [ ] **Algorithm selection (future)**
  - [ ] Create `Algorithm` interface
  - [ ] Implement Kruskal's algorithm
  - [ ] Implement Prim's algorithm
  - [ ] Add `-a, --algorithm` flag for selection
  - [ ] Performance comparison tests

### Path Validation
- [ ] Implement path-finding algorithm to verify maze solvability
- [ ] Add tests to ensure start and end points are accessible
- [ ] Validate maze has no isolated areas

## Phase 4: Enhancement and Polish

### Output Formats
- [ ] **Multiple output formats**
  - [ ] Create `Renderer` interface
  - [ ] Implement Unicode box-drawing renderer
  - [ ] Implement JSON output format
  - [ ] Add `-f, --format` flag
  - [ ] Write tests for each format

### Advanced Features
- [ ] **Start/Goal positioning**
  - [ ] Add `--start` and `--goal` coordinate specification
  - [ ] Visual markers for start/end positions
  - [ ] Validation that positions are valid paths

- [ ] **Solution display**
  - [ ] Implement `--solution` flag
  - [ ] Path-finding algorithm integration
  - [ ] Visual solution path in output

### Performance and Reliability
- [ ] **Error handling improvements**
  - [ ] Comprehensive error messages
  - [ ] Graceful handling of edge cases
  - [ ] Input sanitization

- [ ] **Performance optimization**
  - [ ] Benchmark large maze generation
  - [ ] Memory usage optimization
  - [ ] Concurrent generation for very large mazes

### Testing and Quality
- [ ] **Integration tests**
  - [ ] End-to-end CLI testing
  - [ ] Cross-platform testing
  - [ ] Performance regression tests

- [ ] **Code quality**
  - [ ] Code review and refactoring
  - [ ] Documentation improvements
  - [ ] Static analysis tools integration

## Development Guidelines

### TDD Process
1. **Red**: Write a failing test
2. **Green**: Write minimal code to pass the test
3. **Refactor**: Improve code while keeping tests green
4. Repeat for each feature

### Code Standards
- Follow Go best practices and idioms
- Use meaningful variable and function names
- Write comprehensive comments for public APIs
- Maintain consistent error handling patterns

### Testing Strategy
- Unit tests for all public functions
- Integration tests for CLI features
- Property-based tests for maze validation
- Benchmark tests for performance-critical code

### Git Workflow
- Create feature branches for each phase
- Commit frequently with descriptive messages
- Ensure all tests pass before commits
- Use conventional commit format