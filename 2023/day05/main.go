package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/MueR/adventofcode.go/cast"
	"golang.org/x/sync/errgroup"
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
	seeds, maps := parseInput(input)
	lMin := math.MaxInt
	for _, seed := range seeds {
		for _, m := range maps {
			if dst, contains := m.get(seed); contains {
				seed = dst
			}
		}
		lMin = min(lMin, seed)
	}
	return lMin
}

func part2(input string) int {
	_, maps := parseInput(input)
	seedList := strings.Split(input, "\n\n")[0]
	seeds := parseSeedsRange(seedList)

	m := sync.Mutex{}
	lMin := math.MaxInt
	wg, _ := errgroup.WithContext(context.Background())
	for _, s := range seeds {
		s := s
		wg.Go(func() error {
			l := math.MaxInt
			for i := 0; i < s[1]; i++ {
				num := resolve(s[0]+i, maps)
				l = min(l, num)
			}
			m.Lock()
			lMin = min(lMin, l)
			m.Unlock()
			return nil
		})
	}
	_ = wg.Wait()
	return lMin
}

func parseInput(input string) (seeds []int, maps []Map) {
	parts := strings.Split(input, "\n\n")
	seeds = parseSeeds(parts[0])
	for i := 1; i < len(parts); i++ {
		maps = append(maps, parseMap(strings.Split(parts[i], "\n")))
	}
	return seeds, maps
}

type Range struct {
	From   int
	To     int
	Length int
}

type Map struct {
	ranges []Range
}

func (m *Map) get(v int) (int, bool) {
	l := 0
	r := len(m.ranges) - 1
	for l <= r {
		mid := (l + r) / 2
		rng := m.ranges[mid]
		if v > rng.To {
			l = mid + 1
		} else if v < rng.From {
			r = mid - 1
		} else {
			return v + rng.Length, true
		}
	}
	return 0, false
}

func parseMap(lines []string) Map {
	var ranges []Range
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			// skip header
			continue
		}
		fields := strings.Fields(lines[i])
		dst, _ := strconv.Atoi(fields[0])
		src, _ := strconv.Atoi(fields[1])
		rl, _ := strconv.Atoi(fields[2])
		ranges = append(ranges, Range{
			From:   src,
			To:     src + rl - 1,
			Length: dst - src,
		})
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].From < ranges[j].From
	})
	return Map{ranges: ranges}
}

func parseSeeds(line string) (seeds []int) {
	fields := strings.Fields(strings.TrimSpace(line[6:]))
	for i := 0; i < len(fields); i++ {
		seeds = append(seeds, cast.ToInt(fields[i]))
	}
	return seeds
}

func parseSeedsRange(line string) (seeds [][2]int) {
	fields := strings.Fields(strings.TrimSpace(line[6:]))
	for i := 0; i < len(fields); i += 2 {
		seeds = append(seeds, [2]int{cast.ToInt(fields[i]), cast.ToInt(fields[i+1])})
	}
	return seeds
}

func resolve(v int, maps []Map) int {
	for _, m := range maps {
		if dst, contains := m.get(v); contains {
			v = dst
		}
	}
	return v
}
