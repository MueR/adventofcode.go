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
		ans := part2(input, 1000000)
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1(input string) int {
	grid, rows, cols := parseInput(input)
	var positions []util.Point
	for pos := range grid {
		positions = append(positions, pos)
	}
	res := 0
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			res += dist(rows, cols, positions[i], positions[j], 1)
		}
	}
	return res
}

func part2(input string, emptyReplace int) int {
	board, rows, cols := parseInput(input)

	var positions []util.Point
	for pos := range board {
		positions = append(positions, pos)
	}

	res := 0
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			res += dist(rows, cols, positions[i], positions[j], emptyReplace-1)
		}
	}
	return res
}

func parseInput(input string) (map[util.Point]bool, map[int]bool, map[int]bool) {
	m := make(map[util.Point]bool)
	rows := make(map[int]bool)
	lines := strings.Split(input, "\n")
	for row, line := range lines {
		runes := []rune(line)
		empty := true
		for col, r := range runes {
			if r == '#' {
				m[util.Point{Y: row, X: col}] = true
				empty = false
			}
		}
		if empty {
			rows[row] = true
		}
	}

	cols := make(map[int]bool)
	for col := 0; col < len(lines[0]); col++ {
		empty := true
		for row := 0; row < len(lines); row++ {
			if lines[row][col] == '#' {
				empty = false
				break
			}
		}
		if empty {
			cols[col] = true
		}
	}

	return m, rows, cols
}

func dist(rows map[int]bool, cols map[int]bool, a, b util.Point, empty int) int {
	d := a.Manhattan(b)

	if a.Y < b.Y {
		for row := a.Y; row < b.Y; row++ {
			if rows[row] {
				d += empty
			}
		}
	} else if a.Y > b.Y {
		for row := a.Y - 1; row > b.Y; row-- {
			if rows[row] {
				d += empty
			}
		}
	}

	if a.X < b.X {
		for Col := a.X; Col < b.X; Col++ {
			if cols[Col] {
				d += empty
			}
		}
	} else if a.X > b.X {
		for Col := a.X - 1; Col > b.X; Col-- {
			if cols[Col] {
				d += empty
			}
		}
	}

	return d
}
