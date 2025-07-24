// Package maze provides maze generation and representation functionality.
// This file defines the Renderer interface for different output formats.
package maze

import "fmt"

// Renderer defines the interface for different maze output formats.
type Renderer interface {
	Render(maze *Maze) string
}

// NewRenderer creates a renderer based on the format name.
func NewRenderer(format string) (Renderer, error) {
	switch format {
	case "ascii":
		return &ASCIIRenderer{}, nil
	case "unicode":
		return &UnicodeRenderer{}, nil
	case "json":
		return &JSONRenderer{}, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}

// GetSupportedFormats returns the list of supported output formats.
func GetSupportedFormats() []string {
	return []string{"ascii", "unicode", "json"}
}
