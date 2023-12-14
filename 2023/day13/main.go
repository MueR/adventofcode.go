package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"
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

func part1(input string) (res int) {
	parsed := parseInput(input)
	for _, grid := range parsed {
		res += findCol(grid, 0) + findRow(grid, 0)*100
	}

	return res
}

func part2(input string) (res int) {
	parsed := parseInput(input)
	for _, grid := range parsed {
		res += findCol(grid, 1) + findRow(grid, 1)*100
	}

	return res
}

func parseInput(input string) (ans [][]string) {
	for _, pat := range strings.Split(input, "\n\n") {
		rows := strings.Split(pat, "\n")
		ans = append(ans, rows)
	}
	return ans
}

func findRow(grid []string, maxDiff int) int {
	w, h := len(grid[0]), len(grid)
	for mid := 0; mid < h-1; mid++ {
		diff := 0
		for col := 0; col < w; col++ {
			for off := 0; ; off++ {
				left := mid - off
				right := mid + off + 1
				if left < 0 || right >= h {
					break
				}
				if grid[left][col] != grid[right][col] {
					diff++
					if diff > maxDiff {
						goto skip
					}
				}
			}
		}
		if diff == maxDiff {
			return mid + 1
		}
	skip:
	}
	return 0
}

func findCol(grid []string, maxDiff int) int {
	w, h := len(grid[0]), len(grid)
	for mid := 0; mid < w-1; mid++ {
		diff := 0
		for row := 0; row < h; row++ {
			for off := 0; ; off++ {
				left := mid - off
				right := mid + off + 1
				if left < 0 || right >= w {
					break
				}
				if grid[row][left] != grid[row][right] {
					diff++
					if diff > maxDiff {
						goto skip
					}
				}
			}
		}
		if diff == maxDiff {
			return mid + 1
		}
	skip:
	}
	return 0
}
