package tiltable_grid

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
)

type TiltableGrid struct {
	Grid    [][]byte
	Width   int
	Height  int
	Movable byte
	Blocker byte
	Empty   byte
	seen    []string
}

func NewTilableGrid(width, height int, movable, blocker, empty byte) *TiltableGrid {
	grid := make([][]byte, height)
	for i := range grid {
		grid[i] = make([]byte, width)
	}
	return &TiltableGrid{
		Grid:    grid,
		Width:   width,
		Height:  height,
		Movable: movable,
		Blocker: blocker,
		Empty:   empty,
	}
}

func (g *TiltableGrid) Get(x, y int) byte {
	return g.Grid[y][x]
}

func (g *TiltableGrid) Set(x, y int, b byte) {
	g.Grid[y][x] = b
}

func (g *TiltableGrid) GetRow(y int) []byte {
	return g.Grid[y]
}

func (g *TiltableGrid) SetRow(y int, row []byte) {
	g.Grid[y] = row
}

func (g *TiltableGrid) GetColumn(x int) []byte {
	column := make([]byte, g.Height)
	for y := 0; y < g.Height; y++ {
		column[y] = g.Grid[y][x]
	}
	return column
}

func (g *TiltableGrid) SetColumn(x int, column []byte) {
	for y := 0; y < g.Height; y++ {
		g.Grid[y][x] = column[y]
	}
}

func (g *TiltableGrid) Tilt(dir rune) {
	switch dir {
	case 'N':
		for y := 0; y < g.Height; y++ {
			for x := 0; x < g.Width; x++ {
				g.ApplyPhysics(x, y, dir)
			}
		}
	case 'S':
		for y := g.Height - 1; y >= 0; y-- {
			for x := 0; x < g.Width; x++ {
				g.ApplyPhysics(x, y, dir)
			}
		}
	case 'W':
		for y := 0; y < g.Height; y++ {
			for x := 0; x < g.Width; x++ {
				g.ApplyPhysics(x, y, dir)
			}
		}
	case 'E':
		for y := 0; y < g.Height; y++ {
			for x := g.Width - 1; x >= 0; x-- {
				g.ApplyPhysics(x, y, dir)
			}
		}
	default:
		panic("unknown direction")
	}
}

func (g *TiltableGrid) ApplyPhysics(x, y int, dir rune) {
	object := g.Grid[y][x]
	if object == g.Empty || object == g.Blocker {
		return
	}
	switch dir {
	case 'N':
		if y == 0 {
			return
		}
		g.Grid[y][x] = g.Empty
		for dy := y - 1; dy >= 0; dy-- {
			if g.Grid[dy][x] != g.Empty {
				g.Grid[dy+1][x] = object
				return
			}
		}
		g.Grid[0][x] = object
	case 'S':
		if y == g.Height-1 {
			return
		}
		g.Grid[y][x] = g.Empty
		for dy := y + 1; dy < g.Height; dy++ {
			if g.Grid[dy][x] != g.Empty {
				g.Grid[dy-1][x] = object
				return
			}
		}
		g.Grid[len(g.Grid)-1][x] = object
	case 'W':
		if x == 0 {
			return
		}
		g.Grid[y][x] = g.Empty
		for dx := x - 1; dx >= 0; dx-- {
			if g.Grid[y][dx] != g.Empty {
				g.Grid[y][dx+1] = object
				return
			}
		}
		g.Grid[y][0] = object
	case 'E':
		if x == len(g.Grid[y])-1 {
			return
		}
		g.Grid[y][x] = g.Empty
		for dx := x + 1; dx < g.Width; dx++ {
			if g.Grid[y][dx] != g.Empty {
				g.Grid[y][dx-1] = object
				return
			}
		}
		g.Grid[y][len(g.Grid[y])-1] = object
	default:
		panic("unknown direction")
	}

	return
}

func (g *TiltableGrid) CycleTilt(dirs []rune) {
	for _, dir := range dirs {
		g.Tilt(dir)
	}
}

func (g *TiltableGrid) detectCyclic(seen []string) int {
	for idx := len(seen) - 1; idx >= 0; idx-- {
		for jdx := idx - 1; jdx >= 0; jdx-- {
			if seen[idx] == seen[jdx] {
				return idx - jdx
			}
		}
	}
	return -1
}

func (g *TiltableGrid) hashResult() string {
	h := crc32.New(crc32.IEEETable)
	for _, row := range g.Grid {
		_, _ = h.Write(row)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func (g *TiltableGrid) FindTiltCycle(cycles int, dirs []rune) [][]byte {
	att := 500
	seen := make([]string, 0, att)
	for cycle := 0; cycle < att; cycle++ {
		h := g.hashResult()
		seen = append(seen, h)
		g.CycleTilt(dirs)
	}

	size := g.detectCyclic(seen)
	rem := cycles - att
	for cycle := 0; cycle < rem%size; cycle++ {
		g.CycleTilt(dirs)
	}
	return g.Grid
}

func (g *TiltableGrid) Print() {
	for _, row := range g.Grid {
		fmt.Printf("%s\n", row)
	}
}
