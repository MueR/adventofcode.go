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
	chars := strings.Split(input, ",")
	res := 0
	for _, char := range chars {
		r := hashLabel([]byte(char))
		res += r
	}
	return res
}

func part2(input string) int {
	steps := parseInput(input)
	boxes := map[int][]step{}
	for _, s := range steps {
		currentBox := boxes[s.hash]
		if s.op == '=' {
			boxHasStep := false
			for i, lens := range currentBox {
				if lens.label == s.label {
					currentBox[i] = s
					boxHasStep = true
					break
				}
			}
			if !boxHasStep {
				currentBox = append(currentBox, s)
				boxes[s.hash] = currentBox
				continue
			}
		}
		if s.op == '-' {
			for i, lens := range currentBox {
				if lens.label == s.label {
					currentBox = append(currentBox[:i], currentBox[i+1:]...)
					boxes[s.hash] = currentBox
					break
				}
			}
		}
	}

	res := 0
	for k, v := range boxes {
		for i, s := range v {
			res += (k + 1) * (i + 1) * s.val
		}
	}
	return res
}

func parseInput(input string) []step {
	var steps []step
	for _, str := range strings.Split(input, ",") {
		s, pl := step{}, true
		for i := 0; i < len(str) && pl; i++ {
			char := rune(str[i])
			switch char {
			case '=':
				s.op = char
				s.val = cast.ToInt(str[i+1:])
				pl = false
			case '-':
				s.op = char
				pl = false
			default:
				s.label += string(char)
			}
		}
		s.hash = hashLabel([]byte(s.label))
		steps = append(steps, s)
	}
	return steps
}

func hashLabel(s []byte) (i int) {
	i = 0
	for _, c := range s {
		i = ((i + int(c)) * 17) % 256
	}
	return i
}

type step struct {
	label string
	hash  int
	op    rune
	val   int
}
