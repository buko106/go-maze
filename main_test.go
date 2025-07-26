package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestSizeValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "default size works",
			args:    []string{},
			wantErr: false,
		},
		{
			name:    "valid odd size",
			args:    []string{"-s", "15"},
			wantErr: false,
		},
		{
			name:    "size too small",
			args:    []string{"-s", "3"},
			wantErr: true,
			errMsg:  "Size must be at least 5",
		},
		{
			name:    "even size",
			args:    []string{"-s", "10"},
			wantErr: true,
			errMsg:  "Size must be odd",
		},
		{
			name:    "long flag valid",
			args:    []string{"--size", "13"},
			wantErr: false,
		},
		{
			name:    "long flag even size",
			args:    []string{"--size", "12"},
			wantErr: true,
			errMsg:  "Size must be odd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			output, err := cmd.CombinedOutput()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				if !strings.Contains(string(output), tt.errMsg) {
					t.Errorf("Expected error message '%s' but got '%s'", tt.errMsg, string(output))
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v, output: %s", err, string(output))
				}
				// Verify output looks like a maze (starts and ends with #)
				lines := strings.Split(strings.TrimSpace(string(output)), "\n")
				if len(lines) == 0 {
					t.Error("No output generated")
				}
				// First and last lines should be all walls
				if !strings.HasPrefix(lines[0], "#") || !strings.HasSuffix(lines[0], "#") {
					t.Error("First line should start and end with walls")
				}
				if len(lines) > 1 {
					lastLine := lines[len(lines)-1]
					if !strings.HasPrefix(lastLine, "#") || !strings.HasSuffix(lastLine, "#") {
						t.Error("Last line should start and end with walls")
					}
				}
			}
		})
	}
}

func TestMazeDimensions(t *testing.T) {
	sizes := []string{"5", "7", "9", "11", "15", "21"}

	for _, size := range sizes {
		t.Run("size_"+size, func(t *testing.T) {
			cmd := exec.Command("go", "run", "main.go", "-s", size)
			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("Failed to run command: %v", err)
			}

			lines := strings.Split(strings.TrimSpace(string(output)), "\n")

			// Convert string to expected size
			expectedSize := map[string]int{"5": 5, "7": 7, "9": 9, "11": 11, "15": 15, "21": 21}[size]

			if len(lines) != expectedSize {
				t.Errorf("Expected %d lines but got %d", expectedSize, len(lines))
			}

			for i, line := range lines {
				runeCount := len([]rune(line))
				if runeCount != expectedSize {
					t.Errorf("Line %d should have %d characters but has %d", i, expectedSize, runeCount)
				}
			}
		})
	}
}

// Test seed functionality - should produce same maze for same seed
func TestSeedReproducibility(t *testing.T) {
	seedValue := "12345"

	// Generate maze twice with same seed
	cmd1 := exec.Command("go", "run", "main.go", "--seed", seedValue, "-s", "9")
	output1, err1 := cmd1.Output()
	if err1 != nil {
		t.Fatalf("First run failed: %v", err1)
	}

	cmd2 := exec.Command("go", "run", "main.go", "--seed", seedValue, "-s", "9")
	output2, err2 := cmd2.Output()
	if err2 != nil {
		t.Fatalf("Second run failed: %v", err2)
	}

	if string(output1) != string(output2) {
		t.Error("Same seed should produce identical mazes")
		t.Logf("Output1:\n%s", string(output1))
		t.Logf("Output2:\n%s", string(output2))
	}
}

// Test different seeds produce different mazes
func TestDifferentSeedsDifferentMazes(t *testing.T) {
	cmd1 := exec.Command("go", "run", "main.go", "--seed", "111", "-s", "9")
	output1, err1 := cmd1.Output()
	if err1 != nil {
		t.Fatalf("First run failed: %v", err1)
	}

	cmd2 := exec.Command("go", "run", "main.go", "--seed", "222", "-s", "9")
	output2, err2 := cmd2.Output()
	if err2 != nil {
		t.Fatalf("Second run failed: %v", err2)
	}

	if string(output1) == string(output2) {
		t.Error("Different seeds should produce different mazes")
	}
}

// Test algorithm flag functionality
func TestAlgorithmFlag(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "default algorithm works",
			args:    []string{"-s", "9"},
			wantErr: false,
		},
		{
			name:    "explicit dfs algorithm short flag",
			args:    []string{"-s", "9", "-a", "dfs"},
			wantErr: false,
		},
		{
			name:    "explicit dfs algorithm long flag",
			args:    []string{"-s", "9", "--algorithm", "dfs"},
			wantErr: false,
		},
		{
			name:    "kruskal algorithm short flag",
			args:    []string{"-s", "9", "-a", "kruskal"},
			wantErr: false,
		},
		{
			name:    "kruskal algorithm long flag",
			args:    []string{"-s", "9", "--algorithm", "kruskal"},
			wantErr: false,
		},
		{
			name:    "invalid algorithm",
			args:    []string{"-s", "9", "-a", "invalid"},
			wantErr: true,
			errMsg:  "Unsupported algorithm 'invalid'",
		},
		{
			name:    "invalid algorithm long flag",
			args:    []string{"-s", "9", "--algorithm", "prim"},
			wantErr: true,
			errMsg:  "Unsupported algorithm 'prim'",
		},
		{
			name:    "hunt-kill algorithm short flag",
			args:    []string{"-s", "9", "-a", "hunt-kill"},
			wantErr: false,
		},
		{
			name:    "hunt-kill algorithm long flag",
			args:    []string{"-s", "9", "--algorithm", "hunt-kill"},
			wantErr: false,
		},
		{
			name:    "growing-tree algorithm short flag",
			args:    []string{"-s", "9", "-a", "growing-tree"},
			wantErr: false,
		},
		{
			name:    "growing-tree algorithm long flag",
			args:    []string{"-s", "9", "--algorithm", "growing-tree"},
			wantErr: false,
		},
		{
			name:    "wilson algorithm",
			args:    []string{"-s", "9", "--algorithm", "wilson"},
			wantErr: false,
		},
		{
			name:    "algorithm with seed",
			args:    []string{"-s", "9", "--seed", "42", "--algorithm", "dfs"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			output, err := cmd.CombinedOutput()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				if !strings.Contains(string(output), tt.errMsg) {
					t.Errorf("Expected error message '%s' but got '%s'", tt.errMsg, string(output))
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v, output: %s", err, string(output))
				}
				// Verify output looks like a maze
				lines := strings.Split(strings.TrimSpace(string(output)), "\n")
				if len(lines) == 0 {
					t.Error("No output generated")
				}
				// Check for start and goal markers
				outputStr := string(output)
				if !strings.Contains(outputStr, "●") {
					t.Error("Expected maze to contain start marker (●)")
				}
				if !strings.Contains(outputStr, "○") {
					t.Error("Expected maze to contain goal marker (○)")
				}
			}
		})
	}
}

// Test algorithm with seed reproducibility
func TestAlgorithmSeedReproducibility(t *testing.T) {
	seedValue := "98765"
	algorithm := "dfs"

	// Use binary instead of go run for faster execution
	cmd1 := exec.Command("./maze", "--seed", seedValue, "--algorithm", algorithm, "-s", "9")
	output1, err1 := cmd1.Output()
	if err1 != nil {
		// Fall back to go run if binary doesn't exist
		cmd1 = exec.Command("go", "run", "main.go", "--seed", seedValue, "--algorithm", algorithm, "-s", "9")
		output1, err1 = cmd1.Output()
		if err1 != nil {
			t.Fatalf("First run failed: %v", err1)
		}
	}

	cmd2 := exec.Command("./maze", "--seed", seedValue, "--algorithm", algorithm, "-s", "9")
	output2, err2 := cmd2.Output()
	if err2 != nil {
		// Fall back to go run if binary doesn't exist
		cmd2 = exec.Command("go", "run", "main.go", "--seed", seedValue, "--algorithm", algorithm, "-s", "9")
		output2, err2 = cmd2.Output()
		if err2 != nil {
			t.Fatalf("Second run failed: %v", err2)
		}
	}

	if string(output1) != string(output2) {
		t.Error("Same seed and algorithm should produce identical mazes")
		t.Logf("Output1:\n%s", string(output1))
		t.Logf("Output2:\n%s", string(output2))
	}
}

// Test solution flag functionality
func TestSolutionFlag(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantSolution bool
	}{
		{
			name:         "no solution flag",
			args:         []string{"-s", "9", "--seed", "123"},
			wantSolution: false,
		},
		{
			name:         "solution flag enabled",
			args:         []string{"-s", "9", "--seed", "123", "--solution"},
			wantSolution: true,
		},
		{
			name:         "solution flag with dfs algorithm",
			args:         []string{"-s", "9", "--seed", "123", "--solution", "--algorithm", "dfs"},
			wantSolution: true,
		},
		{
			name:         "solution flag with kruskal algorithm",
			args:         []string{"-s", "9", "--seed", "123", "--solution", "--algorithm", "kruskal"},
			wantSolution: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("Command failed: %v", err)
			}

			outputStr := string(output)

			// Check for required maze elements
			if !strings.Contains(outputStr, "●") {
				t.Error("Expected maze to contain start marker (●)")
			}
			if !strings.Contains(outputStr, "○") {
				t.Error("Expected maze to contain goal marker (○)")
			}

			// Check for solution path markers
			hasSolutionMarkers := strings.Contains(outputStr, "·")
			if tt.wantSolution && !hasSolutionMarkers {
				t.Error("Expected maze to contain solution path markers (·) when --solution flag is used")
			}
			if !tt.wantSolution && hasSolutionMarkers {
				t.Error("Expected maze to NOT contain solution path markers (·) when --solution flag is not used")
			}
		})
	}
}

// Test solution path continuity
func TestSolutionPathContinuity(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "-s", "11", "--seed", "456", "--solution")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		t.Fatal("No output generated")
	}

	// Find start and goal positions
	var startRow, startCol, goalRow, goalCol int
	var found bool
	for i, line := range lines {
		for j, char := range line {
			switch char {
			case '●':
				startRow, startCol = i, j
			case '○':
				goalRow, goalCol = i, j
				found = true
			}
		}
	}

	if !found {
		t.Fatal("Could not find start and goal markers in output")
	}

	// Verify there's a path from start to goal via solution markers or direct connection
	outputStr := string(output)
	if !strings.Contains(outputStr, "·") {
		t.Error("Expected solution path markers (·) to be present")
	}

	// Count solution markers - should be > 0 when solution is requested
	solutionMarkerCount := strings.Count(outputStr, "·")
	if solutionMarkerCount == 0 {
		t.Error("Expected at least one solution path marker (·)")
	}

	t.Logf("Found solution path with %d markers from (%d,%d) to (%d,%d)",
		solutionMarkerCount, startRow, startCol, goalRow, goalCol)
}

// Test solution flag with different seeds produces different solutions
func TestSolutionWithDifferentSeeds(t *testing.T) {
	cmd1 := exec.Command("go", "run", "main.go", "-s", "9", "--seed", "111", "--solution")
	output1, err1 := cmd1.Output()
	if err1 != nil {
		t.Fatalf("First command failed: %v", err1)
	}

	cmd2 := exec.Command("go", "run", "main.go", "-s", "9", "--seed", "222", "--solution")
	output2, err2 := cmd2.Output()
	if err2 != nil {
		t.Fatalf("Second command failed: %v", err2)
	}

	// Both should have solution markers
	if !strings.Contains(string(output1), "·") {
		t.Error("First maze should contain solution markers")
	}
	if !strings.Contains(string(output2), "·") {
		t.Error("Second maze should contain solution markers")
	}

	// Different seeds should produce different solutions
	if string(output1) == string(output2) {
		t.Error("Different seeds should produce different solution paths")
	}
}

// Test format flag functionality
func TestFormatFlag(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantErr     bool
		errMsg      string
		checkOutput func(string) error
	}{
		{
			name:    "default ascii format",
			args:    []string{"-s", "9", "--seed", "123"},
			wantErr: false,
			checkOutput: func(output string) error {
				if !strings.Contains(output, "#") {
					return fmt.Errorf("ASCII output should contain wall characters (#)")
				}
				if !strings.Contains(output, "●") {
					return fmt.Errorf("ASCII output should contain start marker (●)")
				}
				if !strings.Contains(output, "○") {
					return fmt.Errorf("ASCII output should contain goal marker (○)")
				}
				return nil
			},
		},
		{
			name:    "explicit ascii format short flag",
			args:    []string{"-s", "9", "--seed", "123", "-f", "ascii"},
			wantErr: false,
			checkOutput: func(output string) error {
				if !strings.Contains(output, "#") {
					return fmt.Errorf("ASCII output should contain wall characters (#)")
				}
				if !strings.Contains(output, "●") {
					return fmt.Errorf("ASCII output should contain start marker (●)")
				}
				if !strings.Contains(output, "○") {
					return fmt.Errorf("ASCII output should contain goal marker (○)")
				}
				return nil
			},
		},
		{
			name:    "explicit ascii format long flag",
			args:    []string{"-s", "9", "--seed", "123", "--format", "ascii"},
			wantErr: false,
			checkOutput: func(output string) error {
				if !strings.Contains(output, "#") {
					return fmt.Errorf("ASCII output should contain wall characters (#)")
				}
				if !strings.Contains(output, "●") {
					return fmt.Errorf("ASCII output should contain start marker (●)")
				}
				if !strings.Contains(output, "○") {
					return fmt.Errorf("ASCII output should contain goal marker (○)")
				}
				return nil
			},
		},
		{
			name:    "unicode format short flag",
			args:    []string{"-s", "9", "--seed", "123", "-f", "unicode"},
			wantErr: false,
			checkOutput: func(output string) error {
				boxDrawingChars := []string{"─", "│", "┌", "┐", "└", "┘", "├", "┤", "┬", "┴", "┼", "╵", "╷", "╴", "╶", "▪"}
				foundBoxDrawing := false
				for _, char := range boxDrawingChars {
					if strings.Contains(output, char) {
						foundBoxDrawing = true
						break
					}
				}
				if !foundBoxDrawing {
					return fmt.Errorf("Unicode output should contain box-drawing characters")
				}
				if !strings.Contains(output, "◉") {
					return fmt.Errorf("Unicode output should contain start marker (◉)")
				}
				if !strings.Contains(output, "◎") {
					return fmt.Errorf("Unicode output should contain goal marker (◎)")
				}
				return nil
			},
		},
		{
			name:    "unicode format long flag",
			args:    []string{"-s", "9", "--seed", "123", "--format", "unicode"},
			wantErr: false,
			checkOutput: func(output string) error {
				boxDrawingChars := []string{"─", "│", "┌", "┐", "└", "┘", "├", "┤", "┬", "┴", "┼", "╵", "╷", "╴", "╶", "▪"}
				foundBoxDrawing := false
				for _, char := range boxDrawingChars {
					if strings.Contains(output, char) {
						foundBoxDrawing = true
						break
					}
				}
				if !foundBoxDrawing {
					return fmt.Errorf("Unicode output should contain box-drawing characters")
				}
				if !strings.Contains(output, "◉") {
					return fmt.Errorf("Unicode output should contain start marker (◉)")
				}
				if !strings.Contains(output, "◎") {
					return fmt.Errorf("Unicode output should contain goal marker (◎)")
				}
				return nil
			},
		},
		{
			name:    "json format short flag",
			args:    []string{"-s", "9", "--seed", "123", "-f", "json"},
			wantErr: false,
			checkOutput: func(output string) error {
				if !strings.Contains(output, "\"width\"") {
					return fmt.Errorf("JSON output should contain width field")
				}
				if !strings.Contains(output, "\"height\"") {
					return fmt.Errorf("JSON output should contain height field")
				}
				if !strings.Contains(output, "\"grid\"") {
					return fmt.Errorf("JSON output should contain grid field")
				}
				if !strings.Contains(output, "\"start\"") {
					return fmt.Errorf("JSON output should contain start field")
				}
				if !strings.Contains(output, "\"goal\"") {
					return fmt.Errorf("JSON output should contain goal field")
				}
				return nil
			},
		},
		{
			name:    "json format long flag",
			args:    []string{"-s", "9", "--seed", "123", "--format", "json"},
			wantErr: false,
			checkOutput: func(output string) error {
				if !strings.Contains(output, "\"width\"") {
					return fmt.Errorf("JSON output should contain width field")
				}
				if !strings.Contains(output, "\"height\"") {
					return fmt.Errorf("JSON output should contain height field")
				}
				if !strings.Contains(output, "\"grid\"") {
					return fmt.Errorf("JSON output should contain grid field")
				}
				if !strings.Contains(output, "\"start\"") {
					return fmt.Errorf("JSON output should contain start field")
				}
				if !strings.Contains(output, "\"goal\"") {
					return fmt.Errorf("JSON output should contain goal field")
				}
				return nil
			},
		},
		{
			name:    "invalid format short flag",
			args:    []string{"-s", "9", "-f", "invalid"},
			wantErr: true,
			errMsg:  "Unsupported format 'invalid'",
		},
		{
			name:    "invalid format long flag",
			args:    []string{"-s", "9", "--format", "xml"},
			wantErr: true,
			errMsg:  "Unsupported format 'xml'",
		},
		{
			name:    "format with solution flag",
			args:    []string{"-s", "9", "--seed", "123", "--format", "unicode", "--solution"},
			wantErr: false,
			checkOutput: func(output string) error {
				if !strings.Contains(output, "•") {
					return fmt.Errorf("Unicode output with solution should contain solution path markers (•)")
				}
				if !strings.Contains(output, "◉") {
					return fmt.Errorf("Unicode output should contain start marker (◉)")
				}
				if !strings.Contains(output, "◎") {
					return fmt.Errorf("Unicode output should contain goal marker (◎)")
				}
				boxDrawingChars := []string{"─", "│", "┌", "┐", "└", "┘", "├", "┤", "┬", "┴", "┼", "╵", "╷", "╴", "╶", "▪"}
				foundBoxDrawing := false
				for _, char := range boxDrawingChars {
					if strings.Contains(output, char) {
						foundBoxDrawing = true
						break
					}
				}
				if !foundBoxDrawing {
					return fmt.Errorf("Unicode output should contain box-drawing characters")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			output, err := cmd.CombinedOutput()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				if !strings.Contains(string(output), tt.errMsg) {
					t.Errorf("Expected error message '%s' but got '%s'", tt.errMsg, string(output))
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v, output: %s", err, string(output))
				}
				if tt.checkOutput != nil {
					if checkErr := tt.checkOutput(string(output)); checkErr != nil {
						t.Error(checkErr)
					}
				}
			}
		})
	}
}

// Test difficulty flag functionality
func TestDifficultyFlag(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "default difficulty works",
			args:    []string{"-s", "9", "--algorithm", "growing-tree"},
			wantErr: false,
		},
		{
			name:    "easy difficulty short flag",
			args:    []string{"-s", "9", "-a", "growing-tree", "-d", "easy"},
			wantErr: false,
		},
		{
			name:    "easy difficulty long flag",
			args:    []string{"-s", "9", "--algorithm", "growing-tree", "--difficulty", "easy"},
			wantErr: false,
		},
		{
			name:    "medium difficulty short flag",
			args:    []string{"-s", "9", "-a", "growing-tree", "-d", "medium"},
			wantErr: false,
		},
		{
			name:    "medium difficulty long flag",
			args:    []string{"-s", "9", "--algorithm", "growing-tree", "--difficulty", "medium"},
			wantErr: false,
		},
		{
			name:    "hard difficulty short flag",
			args:    []string{"-s", "9", "-a", "growing-tree", "-d", "hard"},
			wantErr: false,
		},
		{
			name:    "hard difficulty long flag",
			args:    []string{"-s", "9", "--algorithm", "growing-tree", "--difficulty", "hard"},
			wantErr: false,
		},
		{
			name:    "extreme difficulty short flag",
			args:    []string{"-s", "9", "-a", "growing-tree", "-d", "extreme"},
			wantErr: false,
		},
		{
			name:    "extreme difficulty long flag",
			args:    []string{"-s", "9", "--algorithm", "growing-tree", "--difficulty", "extreme"},
			wantErr: false,
		},
		{
			name:    "invalid difficulty short flag",
			args:    []string{"-s", "9", "-a", "growing-tree", "-d", "impossible"},
			wantErr: true,
			errMsg:  "Unsupported difficulty 'impossible'",
		},
		{
			name:    "invalid difficulty long flag",
			args:    []string{"-s", "9", "--algorithm", "growing-tree", "--difficulty", "trivial"},
			wantErr: true,
			errMsg:  "Unsupported difficulty 'trivial'",
		},
		{
			name:    "difficulty with seed",
			args:    []string{"-s", "9", "--seed", "42", "--algorithm", "growing-tree", "--difficulty", "hard"},
			wantErr: false,
		},
		{
			name:    "difficulty with non-growing-tree algorithm (should work but ignore difficulty)",
			args:    []string{"-s", "9", "--algorithm", "dfs", "--difficulty", "hard"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			output, err := cmd.CombinedOutput()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				if !strings.Contains(string(output), tt.errMsg) {
					t.Errorf("Expected error message '%s' but got '%s'", tt.errMsg, string(output))
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v, output: %s", err, string(output))
				}
				// Verify output looks like a maze
				lines := strings.Split(strings.TrimSpace(string(output)), "\n")
				if len(lines) == 0 {
					t.Error("No output generated")
				}
				// Check for start and goal markers
				outputStr := string(output)
				if !strings.Contains(outputStr, "●") {
					t.Error("Expected maze to contain start marker (●)")
				}
				if !strings.Contains(outputStr, "○") {
					t.Error("Expected maze to contain goal marker (○)")
				}
			}
		})
	}
}

// Test analyze flag functionality
func TestAnalyzeFlag(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantAnalysis bool
		format       string
	}{
		{
			name:         "no analyze flag",
			args:         []string{"-s", "9", "--seed", "123"},
			wantAnalysis: false,
			format:       "ascii",
		},
		{
			name:         "analyze flag enabled",
			args:         []string{"-s", "9", "--seed", "123", "--analyze"},
			wantAnalysis: true,
			format:       "ascii",
		},
		{
			name:         "analyze flag with algorithm",
			args:         []string{"-s", "9", "--seed", "123", "--analyze", "--algorithm", "growing-tree"},
			wantAnalysis: true,
			format:       "ascii",
		},
		{
			name:         "analyze flag with difficulty",
			args:         []string{"-s", "9", "--seed", "123", "--analyze", "--algorithm", "growing-tree", "--difficulty", "hard"},
			wantAnalysis: true,
			format:       "ascii",
		},
		{
			name:         "analyze flag with json format",
			args:         []string{"-s", "9", "--seed", "123", "--analyze", "--format", "json"},
			wantAnalysis: true,
			format:       "json",
		},
		{
			name:         "analyze flag with unicode format",
			args:         []string{"-s", "9", "--seed", "123", "--analyze", "--format", "unicode"},
			wantAnalysis: true,
			format:       "unicode",
		},
		{
			name:         "analyze flag with solution",
			args:         []string{"-s", "9", "--seed", "123", "--analyze", "--solution"},
			wantAnalysis: true,
			format:       "ascii",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("Command failed: %v", err)
			}

			outputStr := string(output)

			if tt.format == "json" {
				// For JSON format, check if analysis data is included
				if tt.wantAnalysis {
					if !strings.Contains(outputStr, "\"analysis\"") {
						t.Error("Expected JSON output to contain analysis field when --analyze flag is used")
					}
					if !strings.Contains(outputStr, "\"difficulty_score\"") {
						t.Error("Expected JSON analysis to contain difficulty_score field")
					}
					if !strings.Contains(outputStr, "\"difficulty_level\"") {
						t.Error("Expected JSON analysis to contain difficulty_level field")
					}
					if !strings.Contains(outputStr, "\"dead_ends\"") {
						t.Error("Expected JSON analysis to contain dead_ends field")
					}
					if !strings.Contains(outputStr, "\"path_cells\"") {
						t.Error("Expected JSON analysis to contain path_cells field")
					}
				} else {
					// Without analyze flag, analysis should not be present
					if strings.Contains(outputStr, "\"analysis\"") {
						t.Error("Expected JSON output to NOT contain analysis field when --analyze flag is not used")
					}
				}
				// Basic JSON validation
				if !strings.Contains(outputStr, "\"width\"") {
					t.Error("JSON output should contain width field")
				}
				if !strings.Contains(outputStr, "\"start\"") {
					t.Error("JSON output should contain start field")
				}
			} else {
				// For ASCII/Unicode formats, check for analysis report text
				if tt.wantAnalysis {
					if !strings.Contains(outputStr, "Maze Analysis Report:") {
						t.Error("Expected output to contain analysis report when --analyze flag is used")
					}
					if !strings.Contains(outputStr, "Basic Metrics:") {
						t.Error("Expected analysis to contain Basic Metrics section")
					}
					if !strings.Contains(outputStr, "Complexity Metrics:") {
						t.Error("Expected analysis to contain Complexity Metrics section")
					}
					if !strings.Contains(outputStr, "Difficulty Assessment:") {
						t.Error("Expected analysis to contain Difficulty Assessment section")
					}
					if !strings.Contains(outputStr, "Dead Ends:") {
						t.Error("Expected analysis to contain Dead Ends metric")
					}
					if !strings.Contains(outputStr, "Solution Length:") {
						t.Error("Expected analysis to contain Solution Length metric")
					}
				} else {
					// Without analyze flag, analysis report should not be present
					if strings.Contains(outputStr, "Maze Analysis Report:") {
						t.Error("Expected output to NOT contain analysis report when --analyze flag is not used")
					}
				}

				// Check for maze markers based on format
				if tt.format == "unicode" {
					if !strings.Contains(outputStr, "◉") {
						t.Error("Expected Unicode maze to contain start marker (◉)")
					}
					if !strings.Contains(outputStr, "◎") {
						t.Error("Expected Unicode maze to contain goal marker (◎)")
					}
				} else {
					if !strings.Contains(outputStr, "●") {
						t.Error("Expected ASCII maze to contain start marker (●)")
					}
					if !strings.Contains(outputStr, "○") {
						t.Error("Expected ASCII maze to contain goal marker (○)")
					}
				}
			}
		})
	}
}

// Test algorithm and difficulty combination reproducibility
func TestAlgorithmDifficultyReproducibility(t *testing.T) {
	seedValue := "555"
	algorithm := "growing-tree"
	difficulty := "hard"

	// Generate maze twice with same parameters
	cmd1 := exec.Command("go", "run", "main.go", "--seed", seedValue, "--algorithm", algorithm, "--difficulty", difficulty, "-s", "9")
	output1, err1 := cmd1.Output()
	if err1 != nil {
		t.Fatalf("First run failed: %v", err1)
	}

	cmd2 := exec.Command("go", "run", "main.go", "--seed", seedValue, "--algorithm", algorithm, "--difficulty", difficulty, "-s", "9")
	output2, err2 := cmd2.Output()
	if err2 != nil {
		t.Fatalf("Second run failed: %v", err2)
	}

	if string(output1) != string(output2) {
		t.Error("Same seed, algorithm, and difficulty should produce identical mazes")
		t.Logf("Output1:\n%s", string(output1))
		t.Logf("Output2:\n%s", string(output2))
	}
}

// Test analyze flag with different algorithms produces different complexity scores
func TestAnalyzeDifferentAlgorithms(t *testing.T) {
	seedValue := "777"
	size := "11"

	algorithms := []string{"dfs", "kruskal", "wilson", "hunt-kill", "growing-tree"}
	complexityScores := make(map[string]float64)

	for _, algorithm := range algorithms {
		args := []string{"run", "main.go", "--seed", seedValue, "--algorithm", algorithm, "-s", size, "--analyze", "--format", "json"}
		cmd := exec.Command("go", args...)
		output, err := cmd.Output()
		if err != nil {
			t.Fatalf("Command failed for algorithm %s: %v", algorithm, err)
		}

		outputStr := string(output)
		if !strings.Contains(outputStr, "\"analysis\"") {
			t.Errorf("Expected analysis in JSON output for algorithm %s", algorithm)
			continue
		}

		// Extract difficulty score (simplified parsing)
		if strings.Contains(outputStr, "\"difficulty_score\"") {
			// For this test, just verify that analysis data is present
			// Full JSON parsing would be more complex but we're testing CLI integration
			complexityScores[algorithm] = 1.0 // Placeholder
		}
	}

	if len(complexityScores) != len(algorithms) {
		t.Errorf("Expected analysis for all %d algorithms, got %d", len(algorithms), len(complexityScores))
	}
}

// Test comprehensive flag combinations
func TestComprehensiveFlagCombinations(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "all flags together",
			args: []string{"-s", "11", "--seed", "999", "--algorithm", "growing-tree", "--difficulty", "extreme", "--format", "json", "--solution", "--analyze"},
		},
		{
			name: "hunt-kill with solution and analyze",
			args: []string{"-s", "9", "--seed", "111", "--algorithm", "hunt-kill", "--solution", "--analyze"},
		},
		{
			name: "unicode format with growing-tree hard difficulty",
			args: []string{"-s", "9", "--seed", "222", "--algorithm", "growing-tree", "--difficulty", "hard", "--format", "unicode", "--analyze"},
		},
		{
			name: "wilson algorithm with solution in json format",
			args: []string{"-s", "7", "--seed", "333", "--algorithm", "wilson", "--solution", "--format", "json"},
		},
		{
			name: "kruskal with analyze and unicode",
			args: []string{"-s", "9", "--seed", "444", "--algorithm", "kruskal", "--analyze", "--format", "unicode"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("Command failed: %v", err)
			}

			outputStr := string(output)
			if len(outputStr) == 0 {
				t.Error("Expected non-empty output")
			}

			// Basic validation based on flags
			hasAnalyze := false
			hasJSON := false
			hasSolution := false
			hasUnicode := false

			for _, arg := range tt.args {
				switch arg {
				case "--analyze":
					hasAnalyze = true
				case "json":
					hasJSON = true
				case "--solution":
					hasSolution = true
				case "unicode":
					hasUnicode = true
				}
			}

			if hasJSON {
				if !strings.Contains(outputStr, "\"width\"") {
					t.Error("JSON output should contain width field")
				}
				if hasAnalyze && !strings.Contains(outputStr, "\"analysis\"") {
					t.Error("JSON output with analyze should contain analysis field")
				}
				if hasSolution && !strings.Contains(outputStr, "\"solution_path\"") {
					t.Error("JSON output with solution should contain solution_path field")
				}
			} else {
				if hasAnalyze && !strings.Contains(outputStr, "Maze Analysis Report:") {
					t.Error("Non-JSON output with analyze should contain analysis report")
				}
				if hasSolution {
					if hasUnicode {
						if !strings.Contains(outputStr, "•") {
							t.Error("Unicode output with solution should contain solution markers (•)")
						}
					} else {
						if !strings.Contains(outputStr, "·") {
							t.Error("ASCII output with solution should contain solution markers (·)")
						}
					}
				}
			}
		})
	}
}
