package main

import (
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
			expectedSize := len(size) // This is wrong, need to convert string to int

			// Convert string to expected size
			expectedSize = map[string]int{"5": 5, "7": 7, "9": 9, "11": 11, "15": 15, "21": 21}[size]

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
