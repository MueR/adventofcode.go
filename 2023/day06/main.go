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
	races := parseInput(input)

	winWays := 1
	for _, race := range races {
		winWays *= max(1, race.WaysToWin())
	}

	return winWays
}

func part2(input string) int {
	times := strings.Fields(strings.Split(input, "\n")[0])[1:]
	distances := strings.Fields(strings.Split(input, "\n")[1])[1:]
	race := Race{
		Time:     cast.ToInt(strings.Join(times, "")),
		Distance: cast.ToInt(strings.Join(distances, "")),
	}
	return race.WaysToWin()
}

func parseInput(input string) (races []Race) {
	times := strings.Fields(strings.Split(input, "\n")[0])[1:]
	distances := strings.Fields(strings.Split(input, "\n")[1])[1:]
	for i, t := range times {
		races = append(races, Race{
			Time:     cast.ToInt(t),
			Distance: cast.ToInt(distances[i]),
		})
	}
	return races
}

type Race struct {
	Time     int
	Distance int
}

func (r *Race) WaysToWin() int {
	wins := 0
	for t := 1; t < r.Time; t++ {
		d := (r.Time - t) * t
		// fmt.Printf("t: %v, d: %v, %v -- %v\n", t, d, r.Distance, d > r.Distance)
		if d > r.Distance {
			wins++
		}
	}

	return wins
}
