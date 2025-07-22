# Go maze generator Makefile

# Variables
BINARY_NAME=maze
MAIN_PATH=./
BUILD_DIR=./bin
GO_FILES=$(shell find . -name "*.go" -type f)

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build: $(BUILD_DIR)/$(BINARY_NAME)

$(BUILD_DIR)/$(BINARY_NAME): $(GO_FILES)
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# Build and install to current directory
.PHONY: install
install: clean
	go build -o $(BINARY_NAME) $(MAIN_PATH)

# Run tests
.PHONY: test
test:
	go test ./...

# Run tests with verbose output
.PHONY: test-verbose
test-verbose:
	go test -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	go test -cover ./...

# Run tests with detailed coverage report
.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Format code using golangci-lint
.PHONY: fmt
fmt:
	golangci-lint fmt

# Legacy format with go fmt (kept for compatibility)
.PHONY: fmt-go
fmt-go:
	go fmt ./...

# Lint code using golangci-lint
.PHONY: lint
lint:
	golangci-lint run

# Legacy lint with go vet (kept for compatibility)
.PHONY: lint-go
lint-go:
	go vet ./...

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html

# Run the maze generator with default settings
.PHONY: run
run: install
	./$(BINARY_NAME)

# Run examples
.PHONY: examples
examples: install
	@echo "=== Default 21x21 maze ==="
	./$(BINARY_NAME)
	@echo ""
	@echo "=== Small 9x9 maze with seed ==="
	./$(BINARY_NAME) --seed 123 --size 9
	@echo ""
	@echo "=== Large 31x31 maze ==="
	./$(BINARY_NAME) -s 31

# Development workflow
.PHONY: dev
dev: clean fmt lint test build

# Check dependencies
.PHONY: deps
deps:
	go mod tidy
	go mod verify

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build         - Build the maze binary"
	@echo "  install       - Build and install to current directory"
	@echo "  test          - Run all tests"
	@echo "  test-verbose  - Run tests with verbose output"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  coverage      - Generate HTML coverage report"
	@echo "  fmt           - Format code using golangci-lint"
	@echo "  fmt-go        - Format using go fmt (legacy)"
	@echo "  lint          - Run golangci-lint"
	@echo "  lint-go       - Run go vet (legacy)"
	@echo "  clean         - Remove build artifacts"
	@echo "  run           - Build and run with default settings"
	@echo "  examples      - Show example maze outputs"
	@echo "  dev           - Run full development workflow"
	@echo "  deps          - Tidy and verify dependencies"
	@echo "  help          - Show this help message"
