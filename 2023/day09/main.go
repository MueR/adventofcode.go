package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/cast"
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
	parsed := parseInput(input)

	sum := 0
	for _, row := range parsed {
		hist := buildHistory(row)
		hist[len(hist)-1] = append(hist[len(hist)-1], 0)

		for i := len(hist) - 2; i >= 0; i-- {
			hist[i] = append(hist[i], hist[i][len(hist[i])-1]+hist[i+1][len(hist[i+1])-1])
		}
		sum += hist[0][len(hist[0])-1]
	}
	return sum
}

func part2(input string) int {
	parsed := parseInput(input)

	sum := 0
	for _, row := range parsed {
		hist := buildHistory(row)
		hist[len(hist)-1] = append([]int{0}, hist[len(hist)-1]...)
		for i := len(hist) - 2; i >= 0; i-- {
			hist[i] = append([]int{hist[i][0] - hist[i+1][0]}, hist[i]...)
		}
		sum += hist[0][0]
	}
	return sum
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		seq := []int{}
		for _, num := range strings.Fields(line) {
			seq = append(seq, cast.ToInt(num))
		}
		ans = append(ans, seq)
	}
	return ans
}

func buildHistory(seq []int) [][]int {
	res := [][]int{seq}
	for {
		last := len(res) - 1
		done := true
		newRow := make([]int, len(res[last])-1)
		for i := 0; i < len(res[last])-1; i++ {
			newRow[i] = res[last][i+1] - res[last][i]
			if newRow[i] != 0 {
				done = false
			}
		}
		res = append(res, newRow)
		if done {
			return res
		}
	}
}
