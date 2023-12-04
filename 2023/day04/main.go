package main

import (
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
	input string
)

type Card struct {
	Index    int
	WinCount int
	Value    int
	Copies   int
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

func part1(input string) int {
	cards, _ := parseInput(input)

	points := 0
	for _, card := range cards {
		points += card.Value
	}

	return points
}

func part2(input string) int {
	cards, keys := parseInput(input)
	totalCards := 0
	for _, k := range keys {
		if cards[k].WinCount == 0 {
			totalCards += cards[k].Copies
			continue
		}
		for i := k + 1; i <= min(k+cards[k].WinCount, len(cards)); i++ {
			cards[i].Copies += cards[k].Copies
		}
		totalCards += cards[k].Copies
	}
	return totalCards
}

func parseInput(input string) (cards map[int]*Card, keys []int) {
	cards = map[int]*Card{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " | ")
		csf := strings.Split(parts[0], ":")
		cs := strings.Split(strings.TrimSpace(csf[1]), " ")
		cn, _ := strconv.Atoi(strings.TrimSpace(csf[0][strings.Index(parts[0], " "):]))
		ws := strings.Split(strings.TrimSpace(parts[1]), " ")
		winNums := make([]int, len(ws))
		cards[cn] = &Card{
			Index:    cn,
			WinCount: 0,
			Copies:   1,
			Value:    0,
		}
		for _, num := range ws {
			n, err := strconv.Atoi(num)
			if err == nil && n > 0 {
				winNums = append(winNums, n)
			}
		}
		for _, num := range cs {
			n, _ := strconv.Atoi(num)
			if slices.Contains(winNums, n) && n > 0 {
				cards[cn].WinCount++
				cards[cn].Value = max(1, cards[cn].Value*2)
			}
		}
	}
	for k := range cards {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return cards, keys
}
