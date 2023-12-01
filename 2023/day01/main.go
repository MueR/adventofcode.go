package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/maths"
)

//go:embed input.txt
var input string

var wordToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

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

func linesToNumbers(lines []string, mw bool) (numbers []int) {
	for _, line := range lines {
		numbers = append(numbers, (findFirstDigit(line, mw)*10)+findLastDigit(line, mw))
	}
	return numbers
}

func part1(input string) int {
	return maths.SumIntSlice(linesToNumbers(parseInput(input), false))
}

func part2(input string) int {
	return maths.SumIntSlice(linesToNumbers(parseInput(input), true))
}

func findFirstDigit(s string, mw bool) int {
	for i := 0; i < len(s); i++ {
		if isDigit(string(s[i])) {
			d, err := strconv.Atoi(string(s[i]))
			if err == nil {
				return d
			}
		}
		if !mw {
			continue
		}
		for word, digit := range wordToDigit {
			if strings.HasPrefix(s[i:], word) {
				return digit
			}
		}
	}
	return 0
}
func findLastDigit(s string, mw bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if isDigit(string(s[i])) {
			d, err := strconv.Atoi(string(s[i]))
			if err == nil {
				return d
			}
		}
		if !mw {
			continue
		}
		for word, digit := range wordToDigit {
			if strings.HasSuffix(s[:i+1], word) {
				return digit
			}
		}
	}
	return 0
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func parseInput(input string) (ans []string) {
	ans = strings.Split(input, "\n")
	return ans
}
