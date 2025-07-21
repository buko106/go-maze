package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/buko106/go-maze/internal/maze"
)

func main() {
	size := flag.Int("s", 21, "Size of the square maze (must be odd, minimum 5)")
	flag.IntVar(size, "size", 21, "Size of the square maze (must be odd, minimum 5)")
	flag.Parse()

	// Validate size
	if *size < 5 {
		fmt.Fprintf(os.Stderr, "Error: Size must be at least 5, got %d\n", *size)
		os.Exit(1)
	}
	if *size%2 == 0 {
		fmt.Fprintf(os.Stderr, "Error: Size must be odd, got %d\n", *size)
		os.Exit(1)
	}

	generator := maze.NewGenerator()
	m := generator.Generate(*size, *size)
	fmt.Print(m.String())
}
