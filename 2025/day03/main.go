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
	input  string
	parsed [][]int
	solve  solution
)

type (
	solution struct {
		part1  int
		part2  int
		solved bool
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
	parsed = parseInput(input)
	fmt.Printf("Parsed input in %v\n", time.Since(s))
	s = time.Now()
	if part != 2 {
		ans := part1()
		fmt.Printf("Part 1 output: %v  (%v)\n", ans, time.Since(s))
	}
	s = time.Now()
	if part != 1 {
		ans := part2()
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1() (res int64) {
	for _, row := range parsed {
		res += jolts(row, 2)
	}
	return res
}

func part2() (res int64) {
	for _, row := range parsed {
		res += jolts(row, 12)
	}
	return res
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		bank := make([]int, len(line))
		for _, c := range line {
			bank = append(bank, int(c-'0'))
		}
		ans = append(ans, bank)
	}
	return ans
}

func jolts(bank []int, n int) int64 {
	var joltage int64 = 0
	startIdx := 0
	for bc := range n {
		var cur = 0
		for i := startIdx; i <= len(bank)-n+bc; i++ {
			if bank[i] > cur {
				cur = bank[i]
				startIdx = i + 1
			}
		}
		joltage = joltage*10 + int64(cur)
	}
	return joltage
}
