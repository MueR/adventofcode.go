package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	grid "github.com/MueR/adventofcode.go/data-structures/tiltable_grid"
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
		ans := part2(input)
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1(input string) int {
	g := parseInput(input)
	g.Tilt('N')

	return calcScore(g.Grid)
}

func part2(input string) int {
	g := parseInput(input)
	g.FindTiltCycle(1000000000, []rune{'N', 'W', 'S', 'E'})
	return calcScore(g.Grid)
}

func parseInput(input string) (ans *grid.TiltableGrid) {
	g := make([][]byte, 0)
	for _, line := range strings.Split(input, "\n") {
		g = append(g, []byte(line))
	}
	ans = grid.NewTilableGrid(len(g[0]), len(g), 'O', '#', '.')
	ans.Grid = g
	return ans
}

func calcScore(platform [][]byte) (total int) {
	l := len(platform)
	for y, row := range platform {
		for _, object := range row {
			if object == 'O' {
				total += l - y
			}
		}
	}

	return total
}
