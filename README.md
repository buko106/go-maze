# go-maze

A command-line maze generator written in Go.

## Version 0.1.0 (MVP)

Basic maze generation with fixed size output.

### Features

- Generates random 21x21 mazes
- ASCII output using `#` (walls) and ` ` (paths)
- Simple command-line interface

### Installation

```bash
git clone https://github.com/buko106/go-maze.git
cd go-maze
go build -o maze
```

### Usage

```bash
./maze
```

Example output:
```
#######################
#   #       #         #
# # # ##### # ####### #
# #   #     #   #     #
# ##### ##### # # #####
#       #     # #     #
####### # ##### ##### #
#     # #           # #
# ### # ########### # #
#   #             #   #
#######################
```

### Development

```bash
# Run tests
go test ./...

# Build
go build -o maze

# Run
./maze
```

### Testing

This project uses Test-Driven Development (TDD). Run tests frequently during development:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests in verbose mode
go test -v ./...
```

### Architecture

- `main.go`: Entry point and CLI interface
- `internal/maze/`: Core maze generation and representation logic
- Clean separation of concerns for easy testing and extension

### Roadmap

- [ ] Configurable maze size (`-s, --size` option)
- [ ] Seed specification (`--seed` option)
- [ ] Multiple generation algorithms (DFS, Kruskal, Prim)
- [ ] Output format options (Unicode, JSON)
- [ ] Start/goal position specification
- [ ] Solution path display

### Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests first (TDD approach)
4. Implement the feature
5. Ensure all tests pass
6. Submit a pull request

### License

MIT License