package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/util"
)

var (
	//go:embed input.txt
	input  string
	dirMap = map[string]util.Direction{
		"U": util.Up,
		"3": util.Up,
		"D": util.Down,
		"1": util.Down,
		"L": util.Left,
		"2": util.Left,
		"R": util.Right,
		"0": util.Right,
	}
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
	instructions, _ := parseInput(input)
	return getCapacity(instructions)
}

func part2(input string) int {
	_, instructions := parseInput(input)
	return getCapacity(instructions)
}

func parseInput(input string) (p1 []instruction, p2 []instruction) {
	for _, line := range strings.Split(input, "\n") {
		el := strings.Fields(line)
		dir := dirMap[el[0]]
		dist, _ := strconv.Atoi(el[1])
		p1 = append(p1, instruction{dir, dist})

		hDist, _ := strconv.ParseInt(el[2][2:7], 16, 64)
		p2 = append(p2, instruction{dirMap[string(el[2][7])], int(hDist)})
	}
	return p1, p2
}

type instruction struct {
	direction util.Direction
	distance  int
}

func getCapacity(instructions []instruction) int {
	var (
		inside   int
		boundary int
		vertices []util.Position
	)

	start := util.Position{}
	curr := start
	vertices = append(vertices, curr)

	for _, inst := range instructions {
		boundary += inst.distance
		next := curr.Move(inst.direction, int(inst.distance))
		curr = next
		vertices = append(vertices, curr)
	}

	for i := 0; i < len(vertices)-1; i++ {
		a, b := vertices[i], vertices[i+1]
		inside += a.Col*b.Row - a.Row*b.Col
	}

	return boundary/2 + inside/2 + 1
}
