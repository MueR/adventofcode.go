package grid

import (
	"strings"

	"github.com/MueR/adventofcode.go/util"
)

type gridType interface {
	string | int
}

type Grid[T gridType] struct {
	Layout []string
	Map    map[util.Point]T
	Height int
	Width  int
}

func NewGrid[T gridType](layout []string) *Grid[T] {
	points := make(map[util.Point]T)
	for y, line := range layout {
		for x, char := range line {
			points[util.Point{X: x, Y: y}] = T(char)
		}
	}
	return &Grid[T]{
		Layout: layout,
		Height: len(layout),
		Width:  len(layout[0]),
		Map:    points,
	}
}

func (g *Grid[GridType]) Get(x, y int) (b GridType) {
	if y < 0 || y >= len(g.Layout) || x < 0 || x >= len(g.Layout[y]) {
		return
	}
	return GridType(g.Layout[y][x])
}

func (g *Grid[GridType]) Set(x, y int, b byte) {
	if y < 0 || y >= g.Height || x < 0 || x >= g.Width {
		return
	}
	line := []rune(g.Layout[y])
	line[x] = rune(b)
	g.Map[util.Point{X: x, Y: y}] = GridType(b)
	g.Layout[y] = string(line)
}

func (g *Grid[GridType]) Neighbours(x, y int, diagonal bool) (neighbours []GridType) {
	n := [][]int{
		{-1, 0}, {0, -1}, {1, 0}, {0, 1},
	}
	if diagonal {
		n = append(n, []int{-1, -1}, []int{1, -1}, []int{-1, 1}, []int{1, 1})
	}
	for _, d := range n {
		if b := g.Get(x+d[0], y+d[1]); b != nil {
			neighbours = append(neighbours, b)
		}
	}
	return
}

func (g *Grid[GridType]) String() string {
	return strings.Join(g.Layout, "\n")
}
