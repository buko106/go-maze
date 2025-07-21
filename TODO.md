# go-maze Development TODO List

## Phase 1: MVP Implementation (Current Priority)

### Core Functionality
- [ ] **Set up basic project structure**
  - [ ] Create `go.mod` with module `github.com/buko106/go-maze`
  - [ ] Create main package entry point
  - [ ] Set up internal package structure

- [ ] **Implement basic maze generation (TDD)**
  - [ ] Write test for maze creation with fixed dimensions
  - [ ] Write test for maze boundary validation (all edges should be walls)
  - [ ] Write test for maze string representation
  - [ ] Implement `Maze` struct with `Width`, `Height`, `Grid` fields
  - [ ] Implement `Generator` struct with basic random generation
  - [ ] Implement `String()` method for ASCII output

- [ ] **Build and execution verification**
  - [ ] Ensure `go build` produces working binary
  - [ ] Verify maze output is properly formatted
  - [ ] Test on multiple runs to ensure randomness

### Testing Requirements
- [ ] All functions must have unit tests
- [ ] Test coverage should be >80%
- [ ] Tests should follow TDD red-green-refactor cycle

## Phase 2: CLI Features

### Command-line Interface
- [ ] **Size specification feature**
  - [ ] Add CLI argument parsing (consider using `cobra` library)
  - [ ] Implement `-s, --size` flag for square mazes
  - [ ] Add input validation (minimum/maximum size limits)
  - [ ] Write tests for various size inputs
  - [ ] Handle odd/even size requirements for proper maze generation

- [ ] **Seed specification feature**
  - [ ] Implement `--seed` flag for reproducible mazes
  - [ ] Modify Generator to accept seed parameter
  - [ ] Write tests to verify same seed produces same maze
  - [ ] Add seed validation

- [ ] **Help and version display**
  - [ ] Implement `-h, --help` flag
  - [ ] Add `--version` flag
  - [ ] Create comprehensive usage documentation
  - [ ] Add examples in help text

### Input Validation
- [ ] Size range validation (e.g., 5-101, odd numbers only)
- [ ] Seed value validation
- [ ] Error handling for invalid inputs
- [ ] User-friendly error messages

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