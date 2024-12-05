package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
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
	parseInput(input)
	fmt.Printf("Parsed input in %v\n", time.Since(s))
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

func part1(input string) (res int) {
	orderList, pageList := parseInput(input)
	rules := parseRules(orderList)

	for _, list := range strings.Split(pageList, "\n") {
		pages := strings.Split(list, ",")
		if slices.IsSortedFunc(pages, sortFunc(rules)) {
			res += cast.ToInt(pages[len(pages)/2])
		}
	}

	return res
}

func part2(input string) (res int) {
	orderList, pageList := parseInput(input)
	rules := parseRules(orderList)

	for _, upd := range strings.Split(pageList, "\n") {
		p := strings.Split(upd, ",")
		sorted := strings.Split(upd, ",")
		slices.SortStableFunc(sorted, sortFunc(rules))
		if slices.Equal(p, sorted) {
			continue
		}

		res += cast.ToInt(sorted[len(sorted)/2])
	}
	return res
}

func parseInput(input string) (ordering, pages string) {
	ordering, pages, _ = strings.Cut(input, "\n\n")
	return
}

func parseRules(rules string) (res map[int]map[int]struct{}) {
	res = map[int]map[int]struct{}{}
	for _, rule := range strings.Split(rules, "\n") {
		x, y, _ := strings.Cut(rule, "|")
		ix, iy := cast.ToInt(x), cast.ToInt(y)
		if _, ok := res[ix]; !ok {
			res[ix] = make(map[int]struct{})
		}
		res[ix][iy] = struct{}{}
	}
	return res
}

func sortFunc(rules map[int]map[int]struct{}) func(x, y string) int {
	return func(x, y string) int {
		ix, iy := cast.ToInt(x), cast.ToInt(y)
		if _, ok := rules[ix][iy]; ok {
			return -1
		}
		if _, ok := rules[iy][ix]; ok {
			return 1
		}
		return 0
	}
}
