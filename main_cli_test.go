package main

import (
	"os/exec"
	"strings"
	"testing"
)

// Test CLI with default parameters
func TestCLIDefault(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, output)
	}

	// Check that output contains expected maze characters
	outputStr := string(output)
	if !strings.Contains(outputStr, "#") {
		t.Error("Expected maze output to contain wall characters (#)")
	}
	if !strings.Contains(outputStr, "●") {
		t.Error("Expected maze output to contain start marker (●)")
	}
	if !strings.Contains(outputStr, "○") {
		t.Error("Expected maze output to contain goal marker (○)")
	}
}

// Test CLI with size parameter
func TestCLIWithSize(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--size", "7")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, output)
	}

	// Count lines to verify size
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) != 7 {
		t.Errorf("Expected 7 lines for size 7, got %d", len(lines))
	}
}

// Test CLI with seed parameter
func TestCLIWithSeed(t *testing.T) {
	// Run with same seed twice
	cmd1 := exec.Command("go", "run", "main.go", "--seed", "123", "--size", "9")
	output1, err1 := cmd1.CombinedOutput()
	if err1 != nil {
		t.Fatalf("First command failed: %v\nOutput: %s", err1, output1)
	}

	cmd2 := exec.Command("go", "run", "main.go", "--seed", "123", "--size", "9")
	output2, err2 := cmd2.CombinedOutput()
	if err2 != nil {
		t.Fatalf("Second command failed: %v\nOutput: %s", err2, output2)
	}

	if string(output1) != string(output2) {
		t.Error("Same seed should produce identical output")
	}
}

// Test CLI with algorithm parameter
func TestCLIWithAlgorithm(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--algorithm", "dfs", "--size", "7")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, output)
	}

	// Check that output contains expected maze characters
	outputStr := string(output)
	if !strings.Contains(outputStr, "#") {
		t.Error("Expected maze output to contain wall characters (#)")
	}
}

// Test CLI with invalid size (too small)
func TestCLIInvalidSizeSmall(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--size", "3")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Error("Expected command to fail with size 3")
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Error: Size must be at least 5") {
		t.Error("Expected error message about minimum size")
	}
}

// Test CLI with invalid size (even number)
func TestCLIInvalidSizeEven(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--size", "8")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Error("Expected command to fail with even size")
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Error: Size must be odd") {
		t.Error("Expected error message about odd size requirement")
	}
}

// Test CLI with invalid algorithm
func TestCLIInvalidAlgorithm(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--algorithm", "invalid")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Error("Expected command to fail with invalid algorithm")
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Error: Unsupported algorithm") {
		t.Error("Expected error message about unsupported algorithm")
	}
}

// Test CLI short flags
func TestCLIShortFlags(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "-s", "9", "-a", "dfs")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, output)
	}

	// Count lines to verify size
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) != 9 {
		t.Errorf("Expected 9 lines for size 9, got %d", len(lines))
	}
}
