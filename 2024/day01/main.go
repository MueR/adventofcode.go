package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/util"
)

var (
	//go:embed input.txt
	input       string
	left, right []int
	occurrences = make(map[int]int)
)

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
	left, right = parseInput(input)
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
	for i, vl := range left {
		vr := right[i]
		res += int(math.Abs(float64(vr - vl)))
	}

	return res
}

func part2(input string) (res int) {
	for _, vl := range left {
		res += vl * occurrences[vl]
	}
	return res
}

func parseInput(input string) (left, right []int) {
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		l := util.LineToInts(line)
		left = append(left, l[0])
		right = append(right, l[1])
		occurrences[l[1]]++
	}
	sort.Ints(left)
	sort.Ints(right)
	return
}
