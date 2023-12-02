package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

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
		ans := part1(input)
		fmt.Printf("Part 1 output: %v  (%v)\n", ans, time.Since(s))
	}
	s = time.Now()
	if part != 1 {
		ans := part2(input)
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

type pull struct {
	red   int
	green int
	blue  int
}
type game struct {
	num   int
	pulls []pull
	min   pull
}

var games map[int]game

func part1(input string) int {
	possible := 0
	for _, g := range games {
		valid := true
		for _, p := range g.pulls {
			if p.red > 12 || p.green > 12 || p.blue > 12 {
				valid = false
				break
			}
		}
		if valid {
			possible += g.num
		}
	}

	return possible
}

func part2(input string) int {
	possible := 0
	for _, g := range games {
		possible += g.min.red * g.min.green * g.min.blue
	}
	return possible
}

func parseInput(input string) (ans []game) {
	s := strings.Split(input, "\n")
	re := regexp.MustCompile(`Game (\d+): (.*)`)
	games = make(map[int]game)
	for _, line := range s {
		if !re.MatchString(line) {
			panic("bad input")
		}
		g := game{}
		m := re.FindStringSubmatch(line)
		gameNum, _ := strconv.Atoi(m[1])
		g.num = gameNum

		pulls := strings.Split(m[2], ";")
		for _, round := range pulls {
			re := regexp.MustCompile(`(\d+) (\w+)`)
			matches := re.FindAllStringSubmatch(round, -1)
			for _, pd := range matches {
				p := pull{}
				num, _ := strconv.Atoi(pd[1])
				switch pd[2] {
				case "red":
					p.red = num
					g.min.red = int(math.Max(float64(g.min.red), float64(num)))
				case "green":
					p.green = num
					g.min.green = int(math.Max(float64(g.min.green), float64(num)))
				case "blue":
					p.blue = num
					g.min.blue = int(math.Max(float64(g.min.blue), float64(num)))
				}
				g.pulls = append(g.pulls, p)
			}
		}
		games[gameNum] = g
	}
	return ans
}
