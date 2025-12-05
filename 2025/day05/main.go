package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	//go:embed input.txt
	input  string
	ranges []ingredientRange
	solve  solution
	render bool
)

type (
	solution struct {
		part1  int64
		part2  int64
		solved bool
	}

	ingredientRange struct {
		start int64
		end   int64
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
	flag.BoolVar(&render, "render", false, "render puzzle output")
	flag.Parse()

	s := time.Now()
	solve = parseInput(input)
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

func part1() (res int64) {
	return solve.part1
}

func part2() (res int64) {
	return solve.part2
}

func parseInput(input string) (solve solution) {
	sets := strings.Split(input, "\n\n")
	fresh := sets[0]
	ranges = make([]ingredientRange, 0)
	for _, line := range strings.Split(fresh, "\n") {
		r := strings.SplitN(line, "-", 2)
		start, _ := strconv.ParseInt(r[0], 10, 64)
		end, _ := strconv.ParseInt(r[1], 10, 64)
		ranges = append(ranges, ingredientRange{
			start: start,
			end:   end,
		})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	for _, line := range strings.Split(sets[1], "\n") {
		conv, _ := strconv.Atoi(line)
		id := int64(conv)
		if isFresh(id) {
			solve.part1++
		}
	}

	highest := int64(0)
	for _, r := range ranges {
		inc := int64(0)
		if r.end <= highest {
			continue
		}
		if r.end > highest && r.start <= highest {
			inc = r.end - highest
			highest = r.end
		} else {
			inc = r.end - r.start + 1
			highest = r.end
		}
		solve.part2 += inc
	}

	return solve
}

func isFresh(id int64) bool {
	for _, r := range ranges {
		if r.start <= id && id <= r.end {
			return true
		}
		if r.start > id {
			return false
		}
	}
	return false
}
