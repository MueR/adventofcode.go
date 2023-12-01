package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/MueR/adventofcode.go/cast"
	"github.com/MueR/adventofcode.go/maths"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	lines := parseInput(input)
	return maths.SumIntSlice(linesToNumbers(lines))
}

func part2(input string) int {
	rep := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	input = rep.Replace(input)

	lines := parseInput(input)
	return maths.SumIntSlice(linesToNumbers(lines))
}

func linesToNumbers(lines []string) (numbers []int) {
	re := regexp.MustCompile(`[0-9]`)
	for _, line := range lines {
		chars := re.FindAllString(line, -1)
		numbers = append(numbers, cast.ToInt(chars[0]+chars[len(chars)-1]))
	}
	return numbers
}

func parseInput(input string) (ans []string) {
	ans = strings.Split(input, "\n")
	return ans
}
