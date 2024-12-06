package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/MueR/adventofcode.go/data-structures/tilemap"
)

var (
	//go:embed input.txt
	input          string
	reader         *strings.Reader
	parsed         [][]int
	partAPositions []tilemap.Container[rune]
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
	s = time.Now()
	if part != 1 {
		reader.Reset(input)
		ans := part2()
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1() (res int) {
	m := tilemap.FromInput(reader)

	start, _ := m.FirstContainerWith('^')
	dx, dy := 0, -1
	x, y := start.Location()

	for t, ok := m.TileAt(x+dx, y+dy); ok; t, ok = m.TileAt(x+dx, y+dy) {
		if t == '#' {
			dx, dy = rotate(dx, dy)
			continue
		}
		m.SetTile(x, y, 'X')
		x, y = x+dx, y+dy
	}

	partAPositions = m.AllContainersWith('X')
	return len(partAPositions) + 1
}

func part2() (res int) {
	m := tilemap.FromInput(reader)

	var result atomic.Int32
	var wg sync.WaitGroup
	for _, possible := range partAPositions {
		wg.Add(1)
		go func() {
			defer wg.Done()
			x, y := possible.Location()
			if createsLoop(x, y, m) {
				result.Add(1)
			}
		}()
	}

	wg.Wait()
	return int(result.Load())
}

func rotate(dx, dy int) (int, int) {
	switch {
	case dx == 0 && dy == -1:
		return 1, 0
	case dx == 1 && dy == 0:
		return 0, 1
	case dx == 0 && dy == 1:
		return -1, 0
	case dx == -1 && dy == 0:
		return 0, -1
	}

	panic(fmt.Errorf("invalid vector: %d, %d", dx, dy))
}

func createsLoop(cx, cy int, m *tilemap.Map[rune]) bool {
	start, _ := m.FirstContainerWith('^')
	dx, dy := 0, -1
	x, y := start.Location()
	visited := tilemap.Of[rune](m.Size())
	maxSteps := 1000 // probably should find a way to calculate a better default
	for t, ok := m.TileAt(x+dx, y+dy); ok && maxSteps > 0; t, ok = m.TileAt(x+dx, y+dy) {
		if (x+dx == cx && y+dy == cy) || t == '#' {
			dx, dy = rotate(dx, dy)
			continue
		}
		if known, _ := visited.TileAt(x, y); known != 0 {
			maxSteps--
		}

		visited.SetTile(x, y, 'X')
		x, y = x+dx, y+dy
	}
	return maxSteps == 0
}
