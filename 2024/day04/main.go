package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/data-structures/tilemap"
	"github.com/MueR/adventofcode.go/util"
)

var (
	//go:embed input.txt
	input  string
	reader *strings.Reader
	toFind = "XMAS"
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
	reader = strings.NewReader(input)
	fmt.Printf("Parsed input in %v\n", time.Since(s))
	s = time.Now()
	if part != 2 {
		ans := part1()
		fmt.Printf("Part 1 output: %v  (%v)\n", ans, time.Since(s))
	}
	reader.Reset(input)
	s = time.Now()
	if part != 1 {
		reader.Reset(input)
		ans := part2()
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1() (res int) {
	puzzle := tilemap.FromInput(reader)
	for _, p := range puzzle.Values() {
		if v, ok := puzzle.TileAt(p.X, p.Y); !ok || v != 'X' {
			continue
		}
		for _, possible := range from(puzzle, p.X, p.Y) {
			if possible == toFind {
				res++
			}
		}
	}
	return res
}

func part2() (res int) {
	puzzle := tilemap.FromInput(reader)
	for r, p := range puzzle.Values() {
		// Skip if not an 'A'
		if r != 'A' {
			continue
		}

		tl, tlOk := puzzle.TileAt(p.X-1, p.Y-1)
		tr, trOk := puzzle.TileAt(p.X+1, p.Y-1)
		bl, blOk := puzzle.TileAt(p.X-1, p.Y+1)
		br, brOk := puzzle.TileAt(p.X+1, p.Y+1)

		if !(tlOk && trOk && blOk && brOk) {
			// We're on an edge, ignore
			continue
		}

		if ((tl == 'M' && br == 'S') || (tl == 'S' && br == 'M')) &&
			((tr == 'M' && bl == 'S') || (tr == 'S' && bl == 'M')) {
			res++
		}
	}
	return res
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, util.LineToInts(line))
	}
	return ans
}

func from(puzzle *tilemap.Map[rune], x, y int) []string {
	result := make([]string, 0, 8)

	for _, delta := range []struct {
		x int
		y int
	}{
		{1, 1}, {1, 0}, {1, -1}, {0, 1},
		{0, -1}, {-1, 1}, {-1, 0}, {-1, -1},
	} {
		if word := along(puzzle, x, y, delta.x, delta.y); word != "" {
			result = append(result, word)
		}
	}

	return result
}

func along(puzzle *tilemap.Map[rune], x, y, dx, dy int) string {
	var word strings.Builder

	var r rune
	var ok bool

	for i := range len(toFind) {
		r, ok = puzzle.TileAt(x+i*dx, y+i*dy)
		if !ok {
			return ""
		}

		word.WriteRune(r)
	}

	return word.String()
}
