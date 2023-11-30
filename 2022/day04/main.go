package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type pair struct {
	start int
	end   int
}

func main() {
	s := time.Now()
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	ranges := make(map[int][]pair)
	i := 0
	for sc.Scan() {
		elfOne, elfTwo := parsePairs(sc.Text())
		var pairs []pair
		pairs = append(pairs, elfOne)
		pairs = append(pairs, elfTwo)
		ranges[i] = pairs
		i++
	}
	fmt.Printf("Parse: %v\n", time.Since(s))

	s = time.Now()
	fmt.Printf("Part One: %d (%v)\n", partOne(ranges), time.Since(s))
	s = time.Now()
	fmt.Printf("Part Two: %d (%v)\n", partTwo(ranges), time.Since(s))
}

func parseInt(data string) int {
	n, e := strconv.Atoi(data)
	if e != nil {
		panic(fmt.Sprintf("Invalid char %s %v", data, e))
	}
	return n
}

func parsePair(data string) pair {
	numbers := strings.Split(data, "-")
	start, end := parseInt(numbers[0]), parseInt(numbers[1])

	return pair{start, end}
}

func parsePairs(data string) (pair, pair) {
	parts := strings.Split(data, ",")

	return parsePair(parts[0]), parsePair(parts[1])
}

func partOne(ranges map[int][]pair) (result int) {
	for _, elves := range ranges {
		if (elves[0].start >= elves[1].start && elves[0].end <= elves[1].end) ||
			(elves[1].start >= elves[0].start && elves[1].end <= elves[0].end) {
			result++
		}
	}
	return result
}

func partTwo(ranges map[int][]pair) (result int) {
	for _, elves := range ranges {
		if elves[0].start <= elves[1].end && elves[1].start <= elves[0].end {
			result++
		}
	}
	return result
}
