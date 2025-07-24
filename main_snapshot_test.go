package main

import (
	"encoding/json"
	"os/exec"
	"testing"
)

// TestFormatFlagSnapshot tests each format with fixed seed for regression testing
func TestFormatFlagSnapshot(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedOutput string
	}{
		{
			name: "ascii format snapshot with solution",
			args: []string{"-s", "7", "--seed", "42", "--format", "ascii", "--solution"},
			expectedOutput: `#######
#●#···#
#·#·#·#
#···#·#
#####·#
#    ○#
#######
`,
		},
		{
			name: "ascii format snapshot without solution",
			args: []string{"-s", "7", "--seed", "42", "--format", "ascii"},
			expectedOutput: `#######
#●#   #
# # # #
#   # #
##### #
#    ○#
#######
`,
		},
		{
			name: "unicode format snapshot with solution",
			args: []string{"-s", "7", "--seed", "42", "--format", "unicode", "--solution"},
			expectedOutput: `┌─┬───┐
│◉│•••│
│•╵•╷•│
│•••│•│
├───┘•│
│    ◎│
└─────┘
`,
		},
		{
			name: "unicode format snapshot without solution",
			args: []string{"-s", "7", "--seed", "42", "--format", "unicode"},
			expectedOutput: `┌─┬───┐
│◉│   │
│ ╵ ╷ │
│   │ │
├───┘ │
│    ◎│
└─────┘
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			output, err := cmd.Output()
			if err != nil {
				t.Errorf("Command failed: %v", err)
				return
			}

			if string(output) != tt.expectedOutput {
				t.Errorf("Snapshot test failed for %s.\nExpected:\n%s\nGot:\n%s",
					tt.name, tt.expectedOutput, string(output))
			}
		})
	}

	// JSON format snapshot tests (structure validation)
	t.Run("json format snapshot with solution", func(t *testing.T) {
		cmd := exec.Command("go", "run", "main.go", "-s", "7", "--seed", "42", "--format", "json", "--solution")
		output, err := cmd.Output()
		if err != nil {
			t.Errorf("JSON command failed: %v", err)
			return
		}

		// Parse JSON to verify structure
		var parsed map[string]interface{}
		err = json.Unmarshal(output, &parsed)
		if err != nil {
			t.Errorf("Failed to parse JSON output: %v", err)
			return
		}

		// Verify essential fields
		if parsed["width"] != float64(7) {
			t.Errorf("Expected width 7, got %v", parsed["width"])
		}
		if parsed["height"] != float64(7) {
			t.Errorf("Expected height 7, got %v", parsed["height"])
		}

		start, ok := parsed["start"].(map[string]interface{})
		if !ok || start["Row"] != float64(1) || start["Col"] != float64(1) {
			t.Errorf("Expected start position (1,1), got %v", parsed["start"])
		}

		goal, ok := parsed["goal"].(map[string]interface{})
		if !ok || goal["Row"] != float64(5) || goal["Col"] != float64(5) {
			t.Errorf("Expected goal position (5,5), got %v", parsed["goal"])
		}

		solutionPath, ok := parsed["solution_path"].([]interface{})
		if !ok || len(solutionPath) == 0 {
			t.Error("Expected solution_path to be a non-empty array")
		}
	})

	t.Run("json format snapshot without solution", func(t *testing.T) {
		cmd := exec.Command("go", "run", "main.go", "-s", "7", "--seed", "42", "--format", "json")
		output, err := cmd.Output()
		if err != nil {
			t.Errorf("JSON command failed: %v", err)
			return
		}

		// Parse JSON to verify structure
		var parsed map[string]interface{}
		err = json.Unmarshal(output, &parsed)
		if err != nil {
			t.Errorf("Failed to parse JSON output: %v", err)
			return
		}

		// Verify essential fields
		if parsed["width"] != float64(7) {
			t.Errorf("Expected width 7, got %v", parsed["width"])
		}
		if parsed["height"] != float64(7) {
			t.Errorf("Expected height 7, got %v", parsed["height"])
		}

		start, ok := parsed["start"].(map[string]interface{})
		if !ok || start["Row"] != float64(1) || start["Col"] != float64(1) {
			t.Errorf("Expected start position (1,1), got %v", parsed["start"])
		}

		goal, ok := parsed["goal"].(map[string]interface{})
		if !ok || goal["Row"] != float64(5) || goal["Col"] != float64(5) {
			t.Errorf("Expected goal position (5,5), got %v", parsed["goal"])
		}

		// Verify no solution path
		if solutionPath, exists := parsed["solution_path"]; exists {
			if pathArray, ok := solutionPath.([]interface{}); ok && len(pathArray) > 0 {
				t.Errorf("Expected no solution path, but got %d positions", len(pathArray))
			}
		}
	})
}
