package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/util"
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
	_ = parseInput(input)
	fmt.Printf("Parsed input in %v\n", time.Since(s))
	s = time.Now()
	if part != 2 {
		ans := part1(input, 25)
		fmt.Printf("Part 1 output: %v  (%v)\n", ans, time.Since(s))
	}
	s = time.Now()
	if part != 1 {
		ans := part2(input, 75)
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1(input string, steps int) (res int) {
	stones := parseInput(input)
	memory := map[memo]int{}
	for _, s := range stones {
		res += evolve(s, steps, memory)
	}
	return res
}

func part2(input string, steps int) (res int) {
	stones := parseInput(input)
	memory := map[memo]int{}
	for _, s := range stones {
		res += evolve(s, steps, memory)
	}
	return res
}

func parseInput(input string) (ans []int) {
	return util.LineToInts(input)
}

type memo struct {
	steps int
	stone int
}

func evolve(stone int, steps int, mem map[memo]int) (res int) {
	k := memo{steps, stone}
	if v, ok := mem[k]; ok {
		return v
	}
	if steps == 0 {
		res = 1
	} else if stone == 0 {
		res = evolve(1, steps-1, mem)
	} else {
		digits := len(strconv.Itoa(stone))
		if digits%2 == 0 {
			mid := digits / 2
			mod := int(math.Pow(10, float64(mid)))
			r := stone % mod
			l := (stone - r) / mod
			res = evolve(r, steps-1, mem) + evolve(l, steps-1, mem)
		} else {
			res = evolve(stone*2024, steps-1, mem)
		}
	}

	mem[k] = res

	return res
}
