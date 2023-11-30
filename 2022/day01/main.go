package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/MueR/adventofcode.go/cast"
	"github.com/MueR/adventofcode.go/maths"
)

//go:embed input.txt
var input string
var totals []int

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	totals = parseInput(input)
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	return totals[0]
}

func part2() int {
	topThree := 0
	for i := 0; i < 3; i++ {
		topThree += totals[i]
	}
	return topThree
}

func parseInput(input string) (totals []int) {
	var elves [][]int
	for _, group := range strings.Split(input, "\n\n") {
		row := []int{}
		for _, line := range strings.Split(group, "\n") {
			row = append(row, cast.ToInt(line))
		}
		elves = append(elves, row)
	}
	for _, items := range elves {
		totals = append(totals, maths.SumIntSlice(items))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	return totals
}
