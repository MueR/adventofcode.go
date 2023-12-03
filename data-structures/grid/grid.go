package grid

import (
	"strings"
)

type gridType interface {
	string | int
}

type Grid[T gridType] struct {
	Layout []string
	Map    map[Point]T
	Height int
	Width  int
}

type Point struct {
	X int
	Y int
}

func NewGrid[T gridType](layout []string) *Grid[T] {
	points := make(map[Point]T)
	for y, line := range layout {
		for x, char := range line {
			points[Point{X: x, Y: y}] = T(char)
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
	g.Layout[y] = string(line)
}

func (g *Grid[GridType]) Neighbours(x, y int) (neighbours []GridType) {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			neighbours = append(neighbours, g.Get(x+dx, y+dy))
		}
	}
	return
}

func (g *Grid[GridType]) String() string {
	return strings.Join(g.Layout, "\n")
}
