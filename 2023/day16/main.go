package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/util"
)

var (
	//go:embed input.txt
	input string
)

type TileType int

const (
	empty TileType = iota
	mirrorRight
	mirrorLeft
	splitterVertical
	splitterHorizontal
)

type Beam struct {
	pos util.Position
	dir util.Direction
}

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 0, "part 1 or 2")
	flag.Parse()

	s := time.Now()
	if part != 2 {
		ans := part1(input)
		fmt.Printf("Part 1 output: %v  (%v)\n", ans, time.Since(s))
	}
	s = time.Now()
	if part != 1 {
		ans := part2(input)
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1(input string) int {
	board, _, _ := parseInput(input)
	return letThereBeLight(board, Beam{
		pos: util.Position{Row: 0, Col: 0},
		dir: util.Right,
	})
}

func part2(input string) (res int) {
	board, width, height := parseInput(input)
	for row := 0; row < height; row++ {
		res = max(
			res,
			letThereBeLight(board, Beam{
				pos: util.Position{Row: row, Col: 0},
				dir: util.Right,
			}),
			letThereBeLight(board, Beam{
				pos: util.Position{Row: row, Col: width - 1},
				dir: util.Left,
			}),
		)
	}
	for col := 0; col < width; col++ {
		res = max(
			res,
			letThereBeLight(board, Beam{
				pos: util.Position{Row: 0, Col: col},
				dir: util.Down,
			}),
			letThereBeLight(board, Beam{
				pos: util.Position{Row: height - 1, Col: col},
				dir: util.Up,
			}),
		)
	}
	return res
}

func parseInput(input string) (grid map[util.Position]TileType, width, height int) {
	grid = make(map[util.Position]TileType)
	lines := strings.Split(input, "\n")
	for row, line := range lines {
		for col, c := range line {
			pos := util.Position{Row: row, Col: col}
			switch c {
			case '.':
				grid[pos] = empty
			case '/':
				grid[pos] = mirrorRight
			case '\\':
				grid[pos] = mirrorLeft
			case '|':
				grid[pos] = splitterVertical
			case '-':
				grid[pos] = splitterHorizontal
			}
		}
	}
	return grid, len(lines[0]), len(lines)
}

func letThereBeLight(board map[util.Position]TileType, first Beam) int {
	cache := make(map[Beam]struct{})
	energized := make(map[util.Position]struct{})
	q := []Beam{first}

	for len(q) != 0 {
		beam := q[0]
		q = q[1:]

		if _, exists := cache[beam]; exists {
			continue
		}

		t, exists := board[beam.pos]
		if !exists {
			continue
		}
		cache[beam] = struct{}{}
		energized[beam.pos] = struct{}{}

		switch t {
		case empty:
			beam.pos = beam.pos.Move(beam.dir, 1)
			q = append(q, beam)
		case mirrorRight:
			switch beam.dir {
			case util.Left:
				beam.dir = util.Down
			case util.Right:
				beam.dir = util.Up
			case util.Up:
				beam.dir = util.Right
			case util.Down:
				beam.dir = util.Left
			default:
				panic(beam.dir)
			}
			beam.pos = beam.pos.Move(beam.dir, 1)
			q = append(q, beam)
		case mirrorLeft:
			switch beam.dir {
			case util.Left:
				beam.dir = util.Up
			case util.Right:
				beam.dir = util.Down
			case util.Up:
				beam.dir = util.Left
			case util.Down:
				beam.dir = util.Right
			default:
				panic(beam.dir)
			}
			beam.pos = beam.pos.Move(beam.dir, 1)
			q = append(q, beam)
		case splitterVertical:
			switch beam.dir {
			case util.Left, util.Right:
				q = append(q, Beam{
					pos: beam.pos.Move(util.Up, 1),
					dir: util.Up,
				})
				q = append(q, Beam{
					pos: beam.pos.Move(util.Down, 1),
					dir: util.Down,
				})
			case util.Up, util.Down:
				beam.pos = beam.pos.Move(beam.dir, 1)
				q = append(q, beam)
			default:
				panic("unhandled default case")
			}
		case splitterHorizontal:
			switch beam.dir {
			case util.Left, util.Right:
				beam.pos = beam.pos.Move(beam.dir, 1)
				q = append(q, beam)
			case util.Up, util.Down:
				q = append(q, Beam{
					pos: beam.pos.Move(util.Left, 1),
					dir: util.Left,
				})
				q = append(q, Beam{
					pos: beam.pos.Move(util.Right, 1),
					dir: util.Right,
				})
			default:
				panic("unhandled default case")
			}
		}
	}

	return len(energized)
}
