package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/data-structures/tilemap"
	"github.com/MueR/adventofcode.go/util"
	aoc "github.com/teivah/go-aoc"
)

var (
	//go:embed input.txt
	input string
	tiles *tilemap.Map[cell]
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
	parseInput(input)
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

func part1(input string) (res int) {
	antennas := make(map[rune][]aoc.Position)
	board := aoc.ParseBoard(strings.Split(input, "\n"), func(r rune, pos aoc.Position) cell {
		switch r {
		case '.':
			return cell{empty: true}
		default:
			antennas[r] = append(antennas[r], pos)
			return cell{frequency: r}
		}
	})

	antinodes := make(map[aoc.Position]struct{})
	for _, positions := range antennas {
		findAntinodes(board, positions, antinodes)
	}

	return len(antinodes)
}

func part2(input string) (res int) {
	antennas := make(map[rune][]aoc.Position)
	board := aoc.ParseBoard(strings.Split(input, "\n"), func(r rune, pos aoc.Position) cell {
		switch r {
		case '.':
			return cell{empty: true}
		default:
			antennas[r] = append(antennas[r], pos)
			return cell{frequency: r}
		}
	})

	antinodes := make(map[aoc.Position]struct{})
	for _, positions := range antennas {
		if len(positions) <= 1 {
			continue
		}
		findAntinodes2(board, positions, antinodes)
		for _, pos := range positions {
			antinodes[pos] = struct{}{}
		}
	}

	return len(antinodes)
}

func parseInput(input string) {
	antennas := make(map[rune][]util.Point)
	tiles = tilemap.ConvertInputOfPositionAware[cell](strings.NewReader(input), func(r rune, col, row int) cell {
		switch r {
		case '.':
			return cell{empty: true}
		default:
			antennas[r] = append(antennas[r], util.Point{X: col, Y: row})
			return cell{frequency: r}
		}
	})
}

type cell struct {
	empty     bool
	frequency rune
}

func findAntinodes(board aoc.Board[cell], positions []aoc.Position, antinodes map[aoc.Position]struct{}) {
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			x := positions[i]
			y := positions[j]
			setAntinode(board, x, y, antinodes)
			setAntinode(board, y, x, antinodes)
		}
	}
}

func setAntinode(board aoc.Board[cell], x, y aoc.Position, antinodes map[aoc.Position]struct{}) {
	d := delta(x, y)
	z := delta(y, d)
	if board.Contains(z) {
		antinodes[z] = struct{}{}
	}
}

func delta(x, y aoc.Position) aoc.Position {
	return aoc.Position{
		Row: x.Row - y.Row,
		Col: x.Col - y.Col,
	}
}

func visualize(antiNodes map[util.Point]struct{}) {
	mx, my := tiles.Size()
	sb := strings.Builder{}
	for y := 0; y < my; y++ {
		for x := 0; x < mx; x++ {
			if c, ok := tiles.TileAt(x, y); ok {
				if _, ok := antiNodes[util.Point{X: x, Y: y}]; ok {
					sb.WriteRune('#')
				} else if c.empty {
					sb.WriteRune('.')
				} else {
					sb.WriteRune(c.frequency)
				}
			}
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func findAntinodes2(board aoc.Board[cell], positions []aoc.Position, antinodes map[aoc.Position]struct{}) {
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			x := positions[i]
			y := positions[j]
			setAntinode2(board, x, y, antinodes)
			setAntinode2(board, y, x, antinodes)
		}
	}
}

func setAntinode2(board aoc.Board[cell], x, y aoc.Position, antinodes map[aoc.Position]struct{}) {
	d := delta(x, y)
	z := delta(y, d)
	if board.Contains(z) {
		antinodes[z] = struct{}{}
		setAntinode2(board, y, z, antinodes)
	}
}
