package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/cast"
	"github.com/MueR/adventofcode.go/util"
)

var (
	//go:embed input.txt
	input    string
	validMul = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
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

func part1(line string) (res int) {
	for _, in := range validMul.FindAllStringSubmatch(line, -1) {
		res += cast.ToInt(in[1]) * cast.ToInt(in[2])
	}
	return res
}

func part2(line string) (res int) {
	r := strings.NewReader(line)
	for section := range util.SectionsOf(r, "do()") {
		check, _, _ := strings.Cut(section, "don't()")
		res += part1(check)
	}
	return res
}
