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
	layout []string
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
	parseInput(input)
	numbers := 0
	for y, _ := range layout {
		for x := 0; x < len(layout[y]); x++ {
			if layout[y][x] == '.' || util.IsDigit(layout[y][x]) {
				continue
			}
			layout[y] = layout[y][:x] + "." + layout[y][x+1:]
			r := [][]int{
				{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1},
				{x - 1, y + 0} /*{0, 0},*/, {x + 1, y + 0},
				{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
			}
			for _, v := range r {
				n, l := getCurrentNumber(layout, v[0], v[1])
				if n > 0 {
					numbers += n
					layout[v[1]] = l
				}
			}
		}
	}

	return numbers
}

func part2(input string) int {
	parseInput(input)
	numbers := 0
	for y, _ := range layout {
		for x := 0; x < len(layout[y]); x++ {
			if layout[y][x] != '*' {
				continue
			}
			found := []int{}
			layout[y] = layout[y][:x] + "." + layout[y][x+1:]
			r := [][]int{
				{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1},
				{x - 1, y + 0} /*{0, 0},*/, {x + 1, y + 0},
				{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
			}
			for _, v := range r {
				n, l := getCurrentNumber(layout, v[0], v[1])
				if n > 0 {
					found = append(found, n)
					layout[v[1]] = l
				}
			}
			if len(found) == 2 {
				numbers += found[0] * found[1]
			}
		}
	}

	return numbers
}

func parseInput(input string) {
	layout = strings.Split(input, "\n")
}

func getCurrentNumber(layout []string, x, y int) (n int, nl string) {
	if y < 0 || y >= len(layout) || x < 0 || x >= len(layout[y]) {
		return 0, ""
	}
	line := layout[y]
	char := string(line[x])
	line = line[:x] + "." + line[x+1:]
	for i := x - 1; i >= 0; i-- {
		if !util.IsDigit(line[i]) {
			break
		}
		char = string(line[i]) + char
		line = line[:i] + "." + line[i+1:]
	}
	for i := x + 1; i < len(line); i++ {
		if !util.IsDigit(line[i]) {
			break
		}
		char += string(line[i])
		line = line[:i] + "." + line[i+1:]
	}
	n, _ = strconv.Atoi(char)
	return n, string(line)
}
