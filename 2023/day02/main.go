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
	games map[int]game
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
	parseInput(input)
	fmt.Printf("Parsing: %v\n", time.Since(s))
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

func part1() (possible int) {
	rules := hand{red: 12, green: 13, blue: 14}
	for _, g := range games {
		for _, h := range g.hands {
			if !h.Valid(rules) {
				goto skip
			}
		}
		possible += g.num
	skip:
	}

	return possible
}

func part2() (possible int) {
	for _, g := range games {
		possible += g.min.Pow()
	}
	return possible
}

func parseInput(input string) (ans []game) {
	s := strings.Split(input, "\n")
	games = make(map[int]game)
	for i, line := range s {
		games[i+1] = newGame(line)
	}
	return ans
}
