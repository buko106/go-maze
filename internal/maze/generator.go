package maze

import (
	"math/rand"
	"strings"
	"time"
)

type Maze struct {
	Width  int
	Height int
	Grid   [][]bool // true = wall, false = path
}

type Generator struct {
	rand *rand.Rand
}

func NewGenerator() *Generator {
	return &Generator{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *Generator) Generate(width, height int) *Maze {
	// MVP: 簡単な迷路生成（後でアルゴリズム実装）
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
		for j := range grid[i] {
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				grid[i][j] = true // 外周は壁
			} else {
				grid[i][j] = g.rand.Float32() < 0.3 // 30%の確率で壁
			}
		}
	}

	return &Maze{
		Width:  width,
		Height: height,
		Grid:   grid,
	}
}

func (m *Maze) String() string {
	var sb strings.Builder
	for _, row := range m.Grid {
		for _, cell := range row {
			if cell {
				sb.WriteRune('#')
			} else {
				sb.WriteRune(' ')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
