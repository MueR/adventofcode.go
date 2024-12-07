package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/cast"
	"github.com/MueR/adventofcode.go/maths"
	"github.com/MueR/adventofcode.go/util"
)

var (
	//go:embed input.txt
	input  string
	parsed []calibration
)

func add(a, b int) int { return a + b }
func mul(a, b int) int { return a * b }
func cat(a, b int) int { return maths.Concat(a, b) }

type calibration struct {
	result int
	input  []int
}

func (c calibration) solve(ops []func(a, b int) int) int {
	if res := c.findFormula(c.input[0], 1, ops); res != 0 {
		return res
	}
	return 0
}

func (c calibration) findFormula(partial, index int, ops []func(a, b int) int) int {
	if partial > c.result {
		return 0
	}
	if index == len(c.input) {
		if partial == c.result {
			return c.result
		}
		return 0
	}
	for _, op := range ops {
		np := op(partial, c.input[index])
		if res := c.findFormula(np, index+1, ops); res != 0 {
			return res
		}
	}
	return 0
}

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
	for _, c := range parsed {
		res += c.solve([]func(a, b int) int{add, mul})
	}
	return res
}

func part2() (res int) {
	for _, c := range parsed {
		res += c.solve([]func(a, b int) int{add, mul, cat})
	}
	return res
}

func parseInput(input string) (ans []calibration) {
	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n") {
		res, inputs, _ := strings.Cut(line, ":")
		ans = append(ans, calibration{
			result: cast.ToInt(res),
			input:  util.LineToInts(inputs),
		})
	}
	return ans
}
