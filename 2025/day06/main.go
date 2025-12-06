package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/maths"
)

var (
	//go:embed input.txt
	input       string
	parsed      []string
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

	problem struct {
		numbers   []int
		operation int
		result    int
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

	start := time.Now()
	parsed = parseInput(input)
	fmt.Printf("Parsed input in %v\n", time.Since(start))
	s := time.Now()
	if part != 2 {
		ans := part1()
		fmt.Printf("Part 1 output: %v  (%v)\n", ans, time.Since(s))
	}
	s = time.Now()
	if part != 1 {
		ans := part2()
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
	fmt.Printf("Total time: %v\n", time.Since(start))
}

func part1() (res int) {
	r := regexp.MustCompile("[0-9]+|\\*|\\+")
	problems := make([]problem, 0)
	for _, l := range parsed {
		nmbrs := r.FindAllString(l, -1)
		if len(problems) == 0 {
			problems = make([]problem, len(nmbrs))
		}
		for i, nmbr := range nmbrs {
			nmbr = strings.TrimSpace(nmbr)
			switch nmbr {
			case "+":
				problems[i].operation = 1
			case "*":
				problems[i].operation = 2
			default:
				val, _ := strconv.Atoi(nmbr)
				problems[i].numbers = append(problems[i].numbers, val)
			}
		}
	}
	for _, prob := range problems {
		ps := 0
		if prob.operation == 2 {
			ps = 1
		}
		for _, number := range prob.numbers {
			if prob.operation == 1 {
				ps += number
			} else {
				ps *= number
			}
		}
		solve.part1 += ps
	}
	return solve.part1
}

func part2() (res int) {
	p := problem{}
	ml := 0
	for _, l := range parsed {
		ml = max(ml, len(l))
	}
	for x := ml - 1; x >= 0; x-- {
		nmbr := ""
		for y := 0; y < len(parsed); y++ {
			if x >= len(parsed[y]) {
				continue
			}
			switch parsed[y][x] {
			case '+':
				p.operation = 1
			case '*':
				p.operation = 2
			default:
				nmbr += string(parsed[y][x])
			}
		}
		nmbr = strings.TrimSpace(nmbr)
		v, _ := strconv.Atoi(nmbr)
		p.numbers = append(p.numbers, v)
		if p.operation == 0 {
			continue
		}
		if p.operation == 1 {
			p.result = maths.SumSlice(p.numbers)
		} else {
			p.result = maths.MultiplySlice(p.numbers)
		}
		solve.part2 += p.result
		p = problem{}
		x--
	}

	return solve.part2
}

func parseInput(input string) []string {
	parsed = strings.Split(input, "\n")

	return parsed
}
