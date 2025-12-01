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
	input  string
	parsed []instruction
)

type (
	instruction struct {
		dir string
		inc int
	}
	solution struct {
		stops  int
		passes int
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
	solve := process(parsed)
	return solve.stops
}

func part2() (res int) {
	solve := process(parsed)
	return solve.passes
}

func parseInput(input string) (ans []instruction) {
	for _, line := range strings.Split(input, "\n") {
		inst := instruction{
			dir: line[0:1],
			inc: util.LineToInts(line[1:])[0],
		}
		ans = append(ans, inst)
	}
	return ans
}

func process(inst []instruction) (solve solution) {
	dial := 50
	for _, in := range inst {
		switch in.dir {
		case "L":
			if dial > 0 && dial-in.inc <= 0 {
				solve.passes += 1
			}
			dial -= in.inc
		case "R":
			if dial < 0 && dial+in.inc >= 0 {
				solve.passes += 1
			}
			dial += in.inc
		}
		passes := maths.AbsInt(dial / 100)
		solve.passes += passes
		dial %= 100
		if dial == 0 {
			solve.stops += 1
		}
	}
	return solve
}
