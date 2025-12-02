package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/cast"
)

var (
	//go:embed input.txt
	input  string
	parsed []idRange
	solve  solution
)

type (
	idRange struct {
		start int
		end   int
	}
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

func part1() (res int) {
	solve.process(parsed)
	return solve.part1
}

func part2() (res int) {
	solve.process(parsed)
	return solve.part2
}

func parseInput(input string) (ans []idRange) {
	for _, line := range strings.Split(input, ",") {
		parts := strings.Split(line, "-")

		inst := idRange{
			start: cast.ToInt(parts[0]),
			end:   cast.ToInt(parts[1]),
		}
		ans = append(ans, inst)
	}
	return ans
}

func (solve *solution) process(idRanges []idRange) {
	if solve.solved {
		return
	}
	for _, r := range idRanges {
		for i := r.start; i <= r.end; i++ {
			s := strconv.Itoa(i)
			if s[:len(s)/2] == s[len(s)/2:] {
				solve.part1 += i
			}
			if strings.Contains((s + s)[1:len(s+s)-1], s) {
				solve.part2 += i
			}
		}
	}
	solve.solved = true
}
