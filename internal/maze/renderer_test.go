package maze

import (
	"encoding/json"
	"strings"
	"testing"
	"unicode/utf8"
)

// TestNewRenderer tests the renderer factory function
func TestNewRenderer(t *testing.T) {
	tests := []struct {
		name        string
		format      string
		expectError bool
		expectType  string
	}{
		{
			name:        "ASCII renderer",
			format:      "ascii",
			expectError: false,
			expectType:  "*maze.ASCIIRenderer",
		},
		{
			name:        "Unicode renderer",
			format:      "unicode",
			expectError: false,
			expectType:  "*maze.UnicodeRenderer",
		},
		{
			name:        "JSON renderer",
			format:      "json",
			expectError: false,
			expectType:  "*maze.JSONRenderer",
		},
		{
			name:        "Invalid format",
			format:      "invalid",
			expectError: true,
			expectType:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderer, err := NewRenderer(tt.format)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for format '%s', got none", tt.format)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error for format '%s': %v", tt.format, err)
				return
			}

			if renderer == nil {
				t.Errorf("Expected renderer for format '%s', got nil", tt.format)
			}
		})
	}
}

// TestGetSupportedFormats tests the supported formats function
func TestGetSupportedFormats(t *testing.T) {
	formats := GetSupportedFormats()
	expectedFormats := []string{"ascii", "unicode", "json"}

	if len(formats) != len(expectedFormats) {
		t.Errorf("Expected %d formats, got %d", len(expectedFormats), len(formats))
	}

	for _, expected := range expectedFormats {
		found := false
		for _, format := range formats {
			if format == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected format '%s' not found in supported formats", expected)
		}
	}
}

// TestASCIIRenderer tests ASCII output format
func TestASCIIRenderer(t *testing.T) {
	// Create a simple test maze
	maze := &Maze{
		Width:    5,
		Height:   5,
		Grid:     createTestGrid(5, 5),
		StartRow: 1,
		StartCol: 1,
		GoalRow:  3,
		GoalCol:  3,
	}

	renderer := &ASCIIRenderer{}
	output := renderer.Render(maze)

	// Check for expected characters
	if !strings.Contains(output, "#") {
		t.Error("ASCII output should contain wall characters (#)")
	}
	if !strings.Contains(output, "●") {
		t.Error("ASCII output should contain start marker (●)")
	}
	if !strings.Contains(output, "○") {
		t.Error("ASCII output should contain goal marker (○)")
	}

	// Check line count
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 5 {
		t.Errorf("Expected 5 lines in ASCII output, got %d", len(lines))
	}
}

// TestUnicodeRenderer tests Unicode output format
func TestUnicodeRenderer(t *testing.T) {
	// Create a simple test maze
	maze := &Maze{
		Width:    5,
		Height:   5,
		Grid:     createTestGrid(5, 5),
		StartRow: 1,
		StartCol: 1,
		GoalRow:  3,
		GoalCol:  3,
	}

	renderer := &UnicodeRenderer{}
	output := renderer.Render(maze)

	// Check for expected Unicode characters
	boxDrawingChars := []string{"─", "│", "┌", "┐", "└", "┘", "├", "┤", "┬", "┴", "┼", "╵", "╷", "╴", "╶", "▪"}
	foundBoxDrawing := false
	for _, char := range boxDrawingChars {
		if strings.Contains(output, char) {
			foundBoxDrawing = true
			break
		}
	}
	if !foundBoxDrawing {
		t.Error("Unicode output should contain box-drawing characters")
	}

	if !strings.Contains(output, "◉") {
		t.Error("Unicode output should contain start marker (◉)")
	}
	if !strings.Contains(output, "◎") {
		t.Error("Unicode output should contain goal marker (◎)")
	}

	// Check line count
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 5 {
		t.Errorf("Expected 5 lines in Unicode output, got %d", len(lines))
	}
}

// TestJSONRenderer tests JSON output format
func TestJSONRenderer(t *testing.T) {
	// Create a simple test maze
	maze := &Maze{
		Width:    5,
		Height:   5,
		Grid:     createTestGrid(5, 5),
		StartRow: 1,
		StartCol: 1,
		GoalRow:  3,
		GoalCol:  3,
	}

	renderer := &JSONRenderer{}
	output := renderer.Render(maze)

	// Parse JSON to verify structure
	var parsed JSON
	err := json.Unmarshal([]byte(output), &parsed)
	if err != nil {
		t.Errorf("Failed to parse JSON output: %v", err)
		return
	}

	// Verify JSON structure
	if parsed.Width != 5 {
		t.Errorf("Expected width 5 in JSON, got %d", parsed.Width)
	}
	if parsed.Height != 5 {
		t.Errorf("Expected height 5 in JSON, got %d", parsed.Height)
	}
	if parsed.Start.Row != 1 || parsed.Start.Col != 1 {
		t.Errorf("Expected start position (1,1), got (%d,%d)", parsed.Start.Row, parsed.Start.Col)
	}
	if parsed.Goal.Row != 3 || parsed.Goal.Col != 3 {
		t.Errorf("Expected goal position (3,3), got (%d,%d)", parsed.Goal.Row, parsed.Goal.Col)
	}
}

// TestRendererWithSolutionPath tests renderers with solution path
func TestRendererWithSolutionPath(t *testing.T) {
	// Create a maze with solution path
	maze := &Maze{
		Width:    5,
		Height:   5,
		Grid:     createTestGrid(5, 5),
		StartRow: 1,
		StartCol: 1,
		GoalRow:  3,
		GoalCol:  3,
		SolutionPath: []Position{
			{Row: 1, Col: 2},
			{Row: 2, Col: 2},
			{Row: 3, Col: 2},
		},
	}

	// Test ASCII renderer with solution
	asciiRenderer := &ASCIIRenderer{}
	asciiOutput := asciiRenderer.Render(maze)
	if !strings.Contains(asciiOutput, "·") {
		t.Error("ASCII output should contain solution path marker (·)")
	}

	// Test Unicode renderer with solution
	unicodeRenderer := &UnicodeRenderer{}
	unicodeOutput := unicodeRenderer.Render(maze)
	if !strings.Contains(unicodeOutput, "•") {
		t.Error("Unicode output should contain solution path marker (•)")
	}
	// Verify box-drawing characters are present
	boxDrawingChars := []string{"─", "│", "┌", "┐", "└", "┘", "├", "┤", "┬", "┴", "┼", "╵", "╷", "╴", "╶", "▪"}
	foundBoxDrawing := false
	for _, char := range boxDrawingChars {
		if strings.Contains(unicodeOutput, char) {
			foundBoxDrawing = true
			break
		}
	}
	if !foundBoxDrawing {
		t.Error("Unicode output should contain box-drawing characters")
	}

	// Test JSON renderer with solution
	jsonRenderer := &JSONRenderer{}
	jsonOutput := jsonRenderer.Render(maze)
	var parsed JSON
	err := json.Unmarshal([]byte(jsonOutput), &parsed)
	if err != nil {
		t.Errorf("Failed to parse JSON output with solution: %v", err)
		return
	}
	if len(parsed.SolutionPath) != 3 {
		t.Errorf("Expected 3 solution path positions, got %d", len(parsed.SolutionPath))
	}
}

// TestUnicodeSpecificCharacters tests Unicode-specific character validation
func TestUnicodeSpecificCharacters(t *testing.T) {
	// Create test maze with various elements
	maze := &Maze{
		Width:    7,
		Height:   7,
		Grid:     createTestGrid(7, 7),
		StartRow: 1,
		StartCol: 1,
		GoalRow:  5,
		GoalCol:  5,
		SolutionPath: []Position{
			{Row: 1, Col: 2},
			{Row: 1, Col: 3},
			{Row: 2, Col: 3},
			{Row: 3, Col: 3},
			{Row: 4, Col: 3},
			{Row: 5, Col: 3},
			{Row: 5, Col: 4},
		},
	}

	renderer := &UnicodeRenderer{}
	output := renderer.Render(maze)

	// Count specific Unicode characters
	boxDrawingChars := []string{"─", "│", "┌", "┐", "└", "┘", "├", "┤", "┬", "┴", "┼", "╵", "╷", "╴", "╶", "▪"}
	wallCount := 0
	for _, char := range boxDrawingChars {
		wallCount += strings.Count(output, char)
	}
	startCount := strings.Count(output, "◉")
	goalCount := strings.Count(output, "◎")
	solutionCount := strings.Count(output, "•")
	spaceCount := strings.Count(output, " ")

	// Validate character counts
	if wallCount == 0 {
		t.Error("Unicode output should contain box-drawing wall characters")
	}
	if startCount != 1 {
		t.Errorf("Expected exactly 1 start marker (◉), got %d", startCount)
	}
	if goalCount != 1 {
		t.Errorf("Expected exactly 1 goal marker (◎), got %d", goalCount)
	}
	if solutionCount != 7 {
		t.Errorf("Expected exactly 7 solution markers (•), got %d", solutionCount)
	}
	if spaceCount == 0 {
		t.Error("Unicode output should contain path spaces")
	}

	// Verify no ASCII characters are mixed in
	if strings.Contains(output, "#") {
		t.Error("Unicode output should not contain ASCII wall characters (#)")
	}
	if strings.Contains(output, "●") {
		t.Error("Unicode output should not contain ASCII start marker (●)")
	}
	if strings.Contains(output, "○") {
		t.Error("Unicode output should not contain ASCII goal marker (○)")
	}
	if strings.Contains(output, "·") {
		t.Error("Unicode output should not contain ASCII solution marker (·)")
	}
}

// TestUnicodeVsASCIIDifferences tests the differences between Unicode and ASCII output
func TestUnicodeVsASCIIDifferences(t *testing.T) {
	// Create identical test maze
	maze := &Maze{
		Width:    5,
		Height:   5,
		Grid:     createTestGrid(5, 5),
		StartRow: 1,
		StartCol: 1,
		GoalRow:  3,
		GoalCol:  3,
		SolutionPath: []Position{
			{Row: 1, Col: 2},
			{Row: 2, Col: 2},
		},
	}

	// Render with both renderers
	asciiRenderer := &ASCIIRenderer{}
	unicodeRenderer := &UnicodeRenderer{}

	asciiOutput := asciiRenderer.Render(maze)
	unicodeOutput := unicodeRenderer.Render(maze)

	// Outputs should be different
	if asciiOutput == unicodeOutput {
		t.Error("ASCII and Unicode outputs should be different")
	}

	// Both should have same structure (same line count, same line lengths)
	asciiLines := strings.Split(strings.TrimSpace(asciiOutput), "\n")
	unicodeLines := strings.Split(strings.TrimSpace(unicodeOutput), "\n")

	if len(asciiLines) != len(unicodeLines) {
		t.Errorf("ASCII and Unicode outputs should have same line count: ASCII=%d, Unicode=%d",
			len(asciiLines), len(unicodeLines))
	}

	for i := 0; i < len(asciiLines) && i < len(unicodeLines); i++ {
		asciiLen := len([]rune(asciiLines[i]))
		unicodeLen := len([]rune(unicodeLines[i]))
		if asciiLen != unicodeLen {
			t.Errorf("Line %d should have same character count: ASCII=%d, Unicode=%d",
				i, asciiLen, unicodeLen)
		}
	}
}

// TestUnicodeCharacterEncoding tests proper Unicode character encoding
func TestUnicodeCharacterEncoding(t *testing.T) {
	maze := &Maze{
		Width:  3,
		Height: 3,
		Grid: [][]bool{
			{true, true, true},
			{true, false, true},
			{true, true, true},
		},
		StartRow: 1,
		StartCol: 1,
		GoalRow:  1,
		GoalCol:  1,
	}

	renderer := &UnicodeRenderer{}
	output := renderer.Render(maze)

	// Verify proper UTF-8 encoding by checking byte representation
	outputBytes := []byte(output)
	if len(outputBytes) == 0 {
		t.Error("Unicode output should not be empty")
	}

	// Verify the output contains valid UTF-8
	if !utf8.Valid(outputBytes) {
		t.Error("Unicode output should be valid UTF-8")
	}

	// Test specific Unicode code points
	boxDrawingChars := []string{"─", "│", "┌", "┐", "└", "┘", "├", "┤", "┬", "┴", "┼", "╵", "╷", "╴", "╶", "▪"}
	expectedStart := "◉" // U+25C9 FISHEYE

	foundBoxDrawing := false
	for _, char := range boxDrawingChars {
		if strings.Contains(output, char) {
			foundBoxDrawing = true
			break
		}
	}
	if !foundBoxDrawing {
		t.Error("Expected Unicode box-drawing characters")
	}
	if !strings.Contains(output, expectedStart) {
		t.Errorf("Expected Unicode start character %s (U+25C9)", expectedStart)
	}
}

// TestASCIIRendererSnapshot tests ASCII format with fixed output for regression testing
func TestASCIIRendererSnapshot(t *testing.T) {
	// Generate maze with fixed seed for consistent output
	generator, err := NewGeneratorWithSeedAndAlgorithm("42", "dfs")
	if err != nil {
		t.Fatalf("Failed to create generator: %v", err)
	}

	maze := generator.Generate(7, 7)
	maze.SolutionPath = FindPath(maze) // Add solution path

	renderer := &ASCIIRenderer{}
	output := renderer.Render(maze)

	expected := `#######
#●#···#
#·#·#·#
#···#·#
#####·#
#    ○#
#######
`

	if output != expected {
		t.Errorf("ASCII snapshot test failed.\nExpected:\n%s\nGot:\n%s", expected, output)
	}
}

// TestASCIIRendererSnapshotNoSolution tests ASCII format without solution path
func TestASCIIRendererSnapshotNoSolution(t *testing.T) {
	// Generate maze with fixed seed for consistent output
	generator, err := NewGeneratorWithSeedAndAlgorithm("42", "dfs")
	if err != nil {
		t.Fatalf("Failed to create generator: %v", err)
	}

	maze := generator.Generate(7, 7)
	// No solution path added

	renderer := &ASCIIRenderer{}
	output := renderer.Render(maze)

	expected := `#######
#●#   #
# # # #
#   # #
##### #
#    ○#
#######
`

	if output != expected {
		t.Errorf("ASCII snapshot test without solution failed.\nExpected:\n%s\nGot:\n%s", expected, output)
	}
}

// TestUnicodeRendererSnapshot tests Unicode format with fixed output for regression testing
func TestUnicodeRendererSnapshot(t *testing.T) {
	// Generate maze with fixed seed for consistent output
	generator, err := NewGeneratorWithSeedAndAlgorithm("42", "dfs")
	if err != nil {
		t.Fatalf("Failed to create generator: %v", err)
	}

	maze := generator.Generate(7, 7)
	maze.SolutionPath = FindPath(maze) // Add solution path

	renderer := &UnicodeRenderer{}
	output := renderer.Render(maze)

	expected := `┌─┬───┐
│◉│•••│
│•╵•╷•│
│•••│•│
├───┘•│
│    ◎│
└─────┘
`

	if output != expected {
		t.Errorf("Unicode snapshot test failed.\nExpected:\n%s\nGot:\n%s", expected, output)
	}
}

// TestUnicodeRendererSnapshotNoSolution tests Unicode format without solution path
func TestUnicodeRendererSnapshotNoSolution(t *testing.T) {
	// Generate maze with fixed seed for consistent output
	generator, err := NewGeneratorWithSeedAndAlgorithm("42", "dfs")
	if err != nil {
		t.Fatalf("Failed to create generator: %v", err)
	}

	maze := generator.Generate(7, 7)
	// No solution path added

	renderer := &UnicodeRenderer{}
	output := renderer.Render(maze)

	expected := `┌─┬───┐
│◉│   │
│ ╵ ╷ │
│   │ │
├───┘ │
│    ◎│
└─────┘
`

	if output != expected {
		t.Errorf("Unicode snapshot test without solution failed.\nExpected:\n%s\nGot:\n%s", expected, output)
	}
}

// TestJSONRendererSnapshot tests JSON format with fixed output for regression testing
func TestJSONRendererSnapshot(t *testing.T) {
	// Generate maze with fixed seed for consistent output
	generator, err := NewGeneratorWithSeedAndAlgorithm("42", "dfs")
	if err != nil {
		t.Fatalf("Failed to create generator: %v", err)
	}

	maze := generator.Generate(7, 7)
	maze.SolutionPath = FindPath(maze) // Add solution path

	renderer := &JSONRenderer{}
	output := renderer.Render(maze)

	// Parse JSON to verify structure rather than exact string match
	var parsed JSON
	err = json.Unmarshal([]byte(output), &parsed)
	if err != nil {
		t.Errorf("JSON snapshot test failed to parse: %v", err)
		return
	}

	// Verify key properties
	if parsed.Width != 7 {
		t.Errorf("Expected width 7, got %d", parsed.Width)
	}
	if parsed.Height != 7 {
		t.Errorf("Expected height 7, got %d", parsed.Height)
	}
	if parsed.Start.Row != 1 || parsed.Start.Col != 1 {
		t.Errorf("Expected start (1,1), got (%d,%d)", parsed.Start.Row, parsed.Start.Col)
	}
	if parsed.Goal.Row != 5 || parsed.Goal.Col != 5 {
		t.Errorf("Expected goal (5,5), got (%d,%d)", parsed.Goal.Row, parsed.Goal.Col)
	}
	if len(parsed.SolutionPath) == 0 {
		t.Error("Expected solution path to be present")
	}

	// Verify grid structure (should be 7x7 with proper boundaries)
	if len(parsed.Grid) != 7 {
		t.Errorf("Expected 7 rows in grid, got %d", len(parsed.Grid))
	}
	for i, row := range parsed.Grid {
		if len(row) != 7 {
			t.Errorf("Expected 7 columns in row %d, got %d", i, len(row))
		}
		// Check boundaries are walls
		if !row[0] || !row[6] {
			t.Errorf("Row %d should have wall boundaries", i)
		}
	}
	// Check top and bottom boundaries
	for j := 0; j < 7; j++ {
		if !parsed.Grid[0][j] || !parsed.Grid[6][j] {
			t.Errorf("Column %d should have wall boundaries", j)
		}
	}
}

// TestJSONRendererSnapshotNoSolution tests JSON format without solution path
func TestJSONRendererSnapshotNoSolution(t *testing.T) {
	// Generate maze with fixed seed for consistent output
	generator, err := NewGeneratorWithSeedAndAlgorithm("42", "dfs")
	if err != nil {
		t.Fatalf("Failed to create generator: %v", err)
	}

	maze := generator.Generate(7, 7)
	// No solution path added

	renderer := &JSONRenderer{}
	output := renderer.Render(maze)

	// Parse JSON to verify structure rather than exact string match
	var parsed JSON
	err = json.Unmarshal([]byte(output), &parsed)
	if err != nil {
		t.Errorf("JSON snapshot test without solution failed to parse: %v", err)
		return
	}

	// Verify key properties
	if parsed.Width != 7 {
		t.Errorf("Expected width 7, got %d", parsed.Width)
	}
	if parsed.Height != 7 {
		t.Errorf("Expected height 7, got %d", parsed.Height)
	}
	if parsed.Start.Row != 1 || parsed.Start.Col != 1 {
		t.Errorf("Expected start (1,1), got (%d,%d)", parsed.Start.Row, parsed.Start.Col)
	}
	if parsed.Goal.Row != 5 || parsed.Goal.Col != 5 {
		t.Errorf("Expected goal (5,5), got (%d,%d)", parsed.Goal.Row, parsed.Goal.Col)
	}
	if len(parsed.SolutionPath) != 0 {
		t.Errorf("Expected no solution path, but got %d positions", len(parsed.SolutionPath))
	}

	// Verify grid structure (should be 7x7 with proper boundaries)
	if len(parsed.Grid) != 7 {
		t.Errorf("Expected 7 rows in grid, got %d", len(parsed.Grid))
	}
	for i, row := range parsed.Grid {
		if len(row) != 7 {
			t.Errorf("Expected 7 columns in row %d, got %d", i, len(row))
		}
		// Check boundaries are walls
		if !row[0] || !row[6] {
			t.Errorf("Row %d should have wall boundaries", i)
		}
	}
	// Check top and bottom boundaries
	for j := 0; j < 7; j++ {
		if !parsed.Grid[0][j] || !parsed.Grid[6][j] {
			t.Errorf("Column %d should have wall boundaries", j)
		}
	}
}

// Helper function to create a test grid
func createTestGrid(width, height int) [][]bool {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
		for j := range grid[i] {
			// Create walls around borders and some internal walls
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				grid[i][j] = true // Wall
			} else {
				grid[i][j] = false // Path
			}
		}
	}
	return grid
}
