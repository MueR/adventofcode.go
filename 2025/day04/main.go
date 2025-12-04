package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/data-structures/tilemap"
	"github.com/MueR/adventofcode.go/util"
	"github.com/fatih/color"
	"github.com/hekmon/liveterm"
)

var (
	//go:embed input.txt
	input   string
	parsed  *tilemap.Map[rune]
	render  bool
	redFn   func(a ...interface{}) string
	greenFn func(a ...interface{}) string
)

type (
	solution struct {
		part1  int
		part2  int
		solved bool
	}
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
	flag.BoolVar(&render, "render", false, "render puzzle output")
	flag.Parse()

	redFn = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	greenFn = color.New(color.FgGreen).SprintFunc()

	s := time.Now()
	parsed = parseInput(input)
	fmt.Printf("Parsed input in %v\n", time.Since(s))
	s = time.Now()
	if part != 2 {
		ans := part1()
		fmt.Printf("Part 1 output: %v  (%v)\n", ans, time.Since(s))
	}
	s = time.Now()
	if part != 1 {
		ans := part2()
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1() (res int) {
	w, h := parsed.Size()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			c, ok := parsed.ContainerAt(i, j)
			if !ok {
				// container does not exist
				continue
			}
			if c.Value != '@' {
				// no roll of paper
				continue
			}
			adjacentPaper := 0
			for _, n := range parsed.AllNeighbors(i, j) {
				nc, ib := parsed.ContainerAt(n.X, n.Y)
				if ib && nc.Value == '@' {
					adjacentPaper++
					continue
				}
			}
			if adjacentPaper < 4 {
				res++
			}
		}
	}
	return res
}

func part2() (res int) {
	w, h := parsed.Size()
	if render {
		liveterm.RefreshInterval = 100 * time.Millisecond
		liveterm.Output = os.Stdout
		liveterm.SetMultiLinesUpdateFx(func() []string { return printMap(parsed) })
		err := liveterm.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to start liveterm: %s\n", err)
			os.Exit(1)
		}
		liveterm.ForceUpdate()
	}
	for i := 0; i < w*h; i++ {
		removed := removePaper()
		if render {
			liveterm.ForceUpdate()
			time.Sleep(250 * time.Millisecond)
		}
		if removed == 0 {
			break
		}
		res += removed
	}
	if render {
		liveterm.ForceUpdate()
		liveterm.Stop(false)
	}

	return res
}

func removePaper() (removed int) {
	w, h := parsed.Size()
	toRemove := make([]util.Point, 0)
	toBlank := make([]util.Point, 0)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			c, ok := parsed.ContainerAt(i, j)
			if !ok || c.Value == '.' {
				continue
			}
			if c.Value == 'x' {
				toBlank = append(toBlank, util.Point{X: i, Y: j})
				continue
			}
			adjacentPaper := 0
			for _, n := range parsed.AllNeighbors(i, j) {
				nc, ib := parsed.ContainerAt(n.X, n.Y)
				if !ib || nc.Value == '@' {
					adjacentPaper++
					continue
				}
			}
			if adjacentPaper < 4 {
				toRemove = append(toRemove, util.Point{X: i, Y: j})
			}
		}
	}
	for _, n := range toRemove {
		parsed.SetTile(n.X, n.Y, 'x')
	}
	for _, n := range toBlank {
		parsed.SetTile(n.X, n.Y, '.')
	}
	return len(toRemove)
}

func printMap(tm *tilemap.Map[rune]) (output []string) {
	w, h := tm.Size()
	for y := 0; y < h; y++ {
		output = append(output, "")
		for x := 0; x < w; x++ {
			c, ok := tm.ContainerAt(x, y)
			if !ok {
				continue
			}
			switch c.Value {
			case '@':
				output[y] += greenFn(string(c.Value))
			case 'x':
				output[y] += redFn(string(c.Value))
			default:
				output[y] += string(c.Value)
			}
		}
	}
	return output
}

func parseInput(input string) *tilemap.Map[rune] {
	tm := tilemap.FromInput(strings.NewReader(input))
	return tm
}
