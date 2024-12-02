package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/maths"
	"github.com/MueR/adventofcode.go/util"
)

var (
	//go:embed input.txt
	input string
	list  [][]int
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
	list = parseInput(input)
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

func part1() (res int) {
	for _, line := range list {
		if safe(line) {
			res++
		}
	}

	return res
}

func part2() (res int) {
	for _, line := range list {
		if safe(line) {
			res++
			continue
		}
		for i := 0; i < len(line); i++ {
			clone := make([]int, len(line))
			copy(clone, line)
			if safe(append(clone[:i], clone[i+1:]...)) {
				res++
				break
			}
		}
	}
	return res
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		ans = append(ans, util.LineToInts(line))
	}
	return ans
}

func safe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}
	prev := levels[0]
	dir := levels[1] - prev

	if dir == 0 || maths.AbsInt(dir) > 3 {
		return false
	}

	prev = levels[1]
	for i := 2; i < len(levels); i++ {
		v := levels[i]
		d := v - prev
		if maths.Sign(d) != maths.Sign(dir) {
			return false
		}
		if maths.AbsInt(d) > 3 {
			return false
		}
		prev = v
	}
	return true
}
