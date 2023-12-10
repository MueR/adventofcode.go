package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
	"time"
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
	pipeMap := parseInput(input)
	loop := findLoop(pipeMap)

	return len(loop) / 2
}

func part2(input string) int {
	pipeMap := parseInput(input)
	loop := findLoop(pipeMap)
	a := area(pipeMap, loop)

	return a
}

const (
	pipeVertical      = '|'
	pipeHorizontal    = '-'
	pipeBendNorthEast = 'L'
	pipeBendNorthWest = 'J'
	pipeBendSouthWest = '7'
	pipeBendSouthEast = 'F'
	ground            = '.'
	startingPosition  = 'S'
)

type position struct {
	row, col int
}

func parseInput(input string) (pipeMap [][]byte) {
	pipeMap = make([][]byte, len(input))
	for i, line := range strings.Split(input, "\n") {
		pipeMap[i] = []byte(line)
	}

	return pipeMap
}

func findLoop(pipeMap [][]byte) []position {
	start := findStart(pipeMap)
	seen := make([][]bool, len(pipeMap))
	for i := range seen {
		seen[i] = make([]bool, len(pipeMap[i]))
	}
	loop := []position{start}
	seen[start.row][start.col] = true
	for {
		n := findConnections(pipeMap, loop[len(loop)-1])
		if len(n) != 2 {
			panic("not a loop")
		}
		for len(n) > 0 && seen[n[0].row][n[0].col] {
			n = n[1:]
		}
		if len(n) == 0 {
			break
		}
		loop = append(loop, n[0])
		seen[n[0].row][n[0].col] = true
	}
	return loop
}

func findStart(pipeMap [][]byte) position {
	for row := range pipeMap {
		for col := range pipeMap[row] {
			if pipeMap[row][col] == startingPosition {
				return position{row, col}
			}
		}
	}

	return position{-1, -1}
}

func findConnections(pipeMap [][]byte, pos position) []position {
	var neighbors []position

	shape := pipeMap[pos.row][pos.col]

	switch shape {
	case startingPosition:
		neighbors = findStartNeigbours(pipeMap, pos)
	case pipeVertical:
		if pos.row > 0 {
			neighbors = append(neighbors, position{pos.row - 1, pos.col})
		}
		if pos.row < len(pipeMap)-1 {
			neighbors = append(neighbors, position{pos.row + 1, pos.col})
		}
	case pipeHorizontal:
		if pos.col > 0 {
			neighbors = append(neighbors, position{pos.row, pos.col - 1})
		}
		if pos.col < len(pipeMap[pos.row])-1 {
			neighbors = append(neighbors, position{pos.row, pos.col + 1})
		}
	case pipeBendNorthEast:
		if pos.row > 0 {
			neighbors = append(neighbors, position{pos.row - 1, pos.col})
		}
		if pos.col < len(pipeMap[pos.row])-1 {
			neighbors = append(neighbors, position{pos.row, pos.col + 1})
		}
	case pipeBendNorthWest:
		if pos.row > 0 {
			neighbors = append(neighbors, position{pos.row - 1, pos.col})
		}
		if pos.col > 0 {
			neighbors = append(neighbors, position{pos.row, pos.col - 1})
		}
	case pipeBendSouthWest:
		if pos.row < len(pipeMap)-1 {
			neighbors = append(neighbors, position{pos.row + 1, pos.col})
		}
		if pos.col > 0 {
			neighbors = append(neighbors, position{pos.row, pos.col - 1})
		}
	case pipeBendSouthEast:
		if pos.row < len(pipeMap)-1 {
			neighbors = append(neighbors, position{pos.row + 1, pos.col})
		}
		if pos.col < len(pipeMap[pos.row])-1 {
			neighbors = append(neighbors, position{pos.row, pos.col + 1})
		}
	}

	return neighbors
}

func findStartNeigbours(pipeMap [][]byte, pos position) []position {
	var neighbors []position

	if pos.row > 0 && slices.Contains([]byte{pipeVertical, pipeBendSouthEast, pipeBendSouthWest}, pipeMap[pos.row-1][pos.col]) {
		neighbors = append(neighbors, position{pos.row - 1, pos.col})
	}
	if pos.row < len(pipeMap)-1 && slices.Contains([]byte{pipeVertical, pipeBendNorthEast, pipeBendNorthWest}, pipeMap[pos.row+1][pos.col]) {
		neighbors = append(neighbors, position{pos.row + 1, pos.col})
	}
	if pos.col > 0 && slices.Contains([]byte{pipeHorizontal, pipeBendNorthEast, pipeBendSouthEast}, pipeMap[pos.row][pos.col-1]) {
		neighbors = append(neighbors, position{pos.row, pos.col - 1})
	}
	if pos.col < len(pipeMap[pos.row])-1 && slices.Contains([]byte{pipeHorizontal, pipeBendNorthWest, pipeBendSouthWest}, pipeMap[pos.row][pos.col+1]) {
		neighbors = append(neighbors, position{pos.row, pos.col + 1})
	}

	return neighbors
}

func area(pipeMap [][]byte, loop []position) int {
	zoom := make([][]byte, len(pipeMap)*2+1)
	for i := range zoom {
		zoom[i] = make([]byte, len(pipeMap[0])*2+1)
	}

	for row := range zoom {
		for col := range zoom[row] {
			zoom[row][col] = ground
		}
	}

	for i := range loop {
		pos, nextPos := loop[i], loop[(i+1)%len(loop)]

		zoom[pos.row*2+1][pos.col*2+1] = pipeMap[pos.row][pos.col]

		rowDelta, colDelta := nextPos.row-pos.row, nextPos.col-pos.col

		switch {
		case rowDelta == -1 && colDelta == 0:
			zoom[pos.row*2][pos.col*2+1] = pipeVertical
		case rowDelta == 1 && colDelta == 0:
			zoom[pos.row*2+2][pos.col*2+1] = pipeVertical
		case rowDelta == 0 && colDelta == -1:
			zoom[pos.row*2+1][pos.col*2] = pipeHorizontal
		case rowDelta == 0 && colDelta == 1:
			zoom[pos.row*2+1][pos.col*2+2] = pipeHorizontal
		default:
			panic("diagonal pipe")
		}
	}

	outside := 0

	seen := make([][]bool, len(zoom))
	for i := range seen {
		seen[i] = make([]bool, len(zoom[i]))
	}

	var stack []position

	proc := func(pos position) {
		if pos.row < 0 || pos.row >= len(zoom) || pos.col < 0 || pos.col >= len(zoom[pos.row]) {
			return
		}
		if zoom[pos.row][pos.col] != ground {
			return
		}
		if seen[pos.row][pos.col] {
			return
		}

		stack = append(stack, pos)
		seen[pos.row][pos.col] = true
		if pos.row%2 == 1 && pos.col%2 == 1 {
			outside++
		}
	}

	pos := position{0, 0}
	stack = append(stack, pos)
	seen[pos.row][pos.col] = true

	for len(stack) > 0 {
		pos, stack = stack[len(stack)-1], stack[:len(stack)-1]

		proc(position{pos.row - 1, pos.col})
		proc(position{pos.row + 1, pos.col})
		proc(position{pos.row, pos.col - 1})
		proc(position{pos.row, pos.col + 1})
	}

	mapArea := len(pipeMap) * len(pipeMap[0])
	inside := mapArea - outside - len(loop)

	return inside
}
