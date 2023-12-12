package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/util"
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

func part1(input string) (res int) {
	parsed := parseInput(input)
	for _, s := range parsed {
		res += s.Count()
	}
	return res
}

func part2(input string) (res int) {
	parsed := parseInput(input)
	for _, s := range parsed {
		// Yeah, we're expanding the input again... go figure
		dupes := [][]byte(nil)
		sums := []int(nil)
		for i := 0; i < 5; i++ {
			dupes = append(dupes, s.Seq)
			sums = append(sums, s.Sum...)
		}
		s.Seq = bytes.Join(dupes, []byte{'?'})
		s.Sum = sums
		res += s.Count()
	}
	return res
}

func parseInput(input string) (ans []Set) {
	for _, line := range strings.Split(input, "\n") {
		f := strings.Fields(line)
		ans = append(ans, Set{
			Seq: []byte(f[0]),
			Sum: util.ParseIntList(f[1], ","),
		})
	}
	return ans
}

type Set struct {
	Seq []byte
	Sum []int
}

func (s Set) Count() int {
	return Set{
		Seq: append(s.Seq, '.'),
		Sum: s.Sum,
	}.calculate(0, 0, 0, make(map[[3]int]int))
}

func (s Set) calculate(i, sum, group int, mem map[[3]int]int) (cnt int) {
	key := [3]int{i, sum, group}
	val, exist := mem[key]
	if exist {
		return val
	}
	// required for memoization
	defer func() {
		if !exist {
			mem[key] = cnt
		}
	}()
	switch {
	case i == len(s.Seq) && group == len(s.Sum): // end of sequence and last group
		return 1
	case i == len(s.Seq) && group != len(s.Sum): // end of sequence but not last group
		return 0
	case s.Seq[i] == '#': // damaged
		return s.calculate(i+1, sum+1, group, mem)
	case s.Seq[i] == '.' || group == len(s.Sum):
		// operational and last group
		if group < len(s.Sum) && sum == s.Sum[group] {
			// found a match, calculate next
			return s.calculate(i+1, 0, group+1, mem)
		} else if sum == 0 {
			return s.calculate(i+1, 0, group, mem)
		} else {
			return 0
		}
	default:
		// oc = operational, dc = damaged
		oc, dc := 0, s.calculate(i+1, sum+1, group, mem)
		if sum == s.Sum[group] {
			// group complete, calculate next
			oc = s.calculate(i+1, 0, group+1, mem)
		} else if sum == 0 {
			oc = s.calculate(i+1, 0, group, mem)
		}
		return dc + oc
	}
}
