package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/util"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
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
	grid, w, h := parseInput(input)

	return findLeastLossPath(grid, util.Position{Col: w - 1, Row: h - 1}, 0, 3)
}

func part2(input string) int {
	grid, w, h := parseInput(input)

	return findLeastLossPath(grid, util.Position{Col: w - 1, Row: h - 1}, 4, 10)
}

func parseInput(input string) (grid map[util.Position]int, width, height int) {
	lines := strings.Split(input, "\n")
	grid = make(map[util.Position]int)
	for row, line := range lines {
		for col, char := range line {
			grid[util.Position{Col: col, Row: row}] = int(char - '0')
		}
	}
	return grid, len(lines[0]), len(lines)
}

func findLeastLossPath(grid map[util.Position]int, target util.Position, minStraight, maxStraight int) (loss int) {
	type state struct {
		location  util.Location
		sameSteps int
	}
	type queueItem struct {
		state state
		loss  int
	}

	visited := make(map[state]int)
	q := pq.NewWith(func(a, b interface{}) int {
		return a.(queueItem).loss - b.(queueItem).loss
	})

	q.Enqueue(queueItem{
		state: state{
			location:  util.NewLocation(0, 1, util.Right),
			sameSteps: 1,
		},
	})
	q.Enqueue(queueItem{
		state: state{
			location:  util.NewLocation(1, 0, util.Down),
			sameSteps: 1,
		},
	})

	for !q.Empty() {
		tile, _ := q.Dequeue()
		entry := tile.(queueItem)
		pos := entry.state.location.Pos
		if _, exists := grid[pos]; !exists {
			continue
		}
		heat := grid[pos] + entry.loss
		if pos == target {
			return heat
		}
		if vis, exists := visited[entry.state]; exists {
			if vis <= heat {
				continue
			}
		}
		visited[entry.state] = heat
		if entry.state.sameSteps >= minStraight {
			q.Enqueue(queueItem{
				state: state{
					location:  entry.state.location.Turn(util.Left, 1),
					sameSteps: 1,
				},
				loss: heat,
			})
			q.Enqueue(queueItem{
				state: state{
					location:  entry.state.location.Turn(util.Right, 1),
					sameSteps: 1,
				},
				loss: heat,
			})
		}
		if entry.state.sameSteps < maxStraight {
			q.Enqueue(queueItem{
				state: state{
					location:  entry.state.location.Straight(1),
					sameSteps: entry.state.sameSteps + 1,
				},
				loss: heat,
			})
		}
	}

	panic("no path found")
}
