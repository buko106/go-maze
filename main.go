package main

import (
	"fmt"

	"github.com/buko106/go-maze/internal/maze"
)

func main() {
	generator := maze.NewGenerator()
	m := generator.Generate(21, 21)
	fmt.Print(m.String())
}
