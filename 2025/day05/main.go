package main

import (
	"cmp"
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	//go:embed input.txt
	input       string
	ranges      []ingredientRange
	ingredients sort.IntSlice
	solve       solution
	render      bool
)

type (
	solution struct {
		part1  int
		part2  int
		solved bool
	}

	ingredientRange struct {
		start int
		end   int
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

func part1() (res int) {
	cur := 0
	for _, id := range ingredients {
		for ranges[cur].end < id {
			cur++
			if cur >= len(ranges) {
				return solve.part1
			}
		}
		if ranges[cur].start <= id {
			solve.part1++
		}
	}
	return solve.part1
}

func part2() (res int) {
	highest := 0
	for _, r := range ranges {
		inc := 0
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
	return solve.part2
}

func parseInput(input string) (solve solution) {
	sets := strings.Split(input, "\n\n")
	fresh := sets[0]
	ranges = make([]ingredientRange, 0)
	for _, line := range strings.Split(fresh, "\n") {
		r := strings.SplitN(line, "-", 2)
		start, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		ranges = append(ranges, ingredientRange{
			start: start,
			end:   end,
		})
	}

	slices.SortFunc(ranges, func(a ingredientRange, b ingredientRange) int {
		if n := cmp.Compare(a.start, b.start); n != 0 {
			return n
		}
		return cmp.Compare(a.end, b.end)
	})

	for _, line := range strings.Split(sets[1], "\n") {
		num, _ := strconv.Atoi(line)
		ingredients = append(ingredients, num)
	}
	ingredients.Sort()

	return solve
}
