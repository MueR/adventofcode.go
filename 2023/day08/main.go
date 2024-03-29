package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/maths"
)

var (
	//go:embed input.txt
	input string
	ops   []string
	nodes map[string]map[string]string
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
	steps := 0
	current := "AAA"
	ol := len(ops)
	for current != "ZZZ" {
		current = nodes[current][ops[steps%ol]]
		steps++
	}

	return steps
}

func part2(input string) int {
	parseInput(input)
	startNodes := make([]string, 0)
	for k, _ := range nodes {
		if k[2:] == "A" {
			startNodes = append(startNodes, k)
		}
	}
	ol := len(ops)
	allSteps := make([]int, 0)
	for _, startNode := range startNodes {
		steps := 0
		current := startNode
		for current[2:] != "Z" {
			current = nodes[current][ops[steps%ol]]
			steps++
		}
		allSteps = append(allSteps, steps)
	}
	if len(allSteps) == 1 {
		return allSteps[0]
	}
	return maths.LCM(allSteps[0], allSteps[1], allSteps[2:]...)
}

func parseInput(input string) {
	parts := strings.Split(input, "\n\n")
	nodes = make(map[string]map[string]string)
	ops = make([]string, 0)
	ops = strings.Split(parts[0], "")
	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	for _, line := range strings.Split(parts[1], "\n") {
		matches := re.FindStringSubmatch(line)
		nodes[matches[1]] = map[string]string{
			"L": matches[2],
			"R": matches[3],
		}
	}
}

type Instruction struct {
	Ops []string
}
