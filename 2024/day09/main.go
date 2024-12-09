package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	disk := parseInput(input)
	fmt.Printf("Parsed input in %v\n", time.Since(s))
	s = time.Now()
	if part != 2 {
		disk = parseInput(input)
		ans := part1(disk)
		fmt.Printf("Part 1 output: %v  (%v)\n", ans, time.Since(s))
	}
	s = time.Now()
	if part != 1 {
		disk = parseInput(input)
		ans := part2(disk)
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1(disk []cell) (res int) {
	l := 0
	r := len(disk) - 1
	for l < r {
		if disk[r].empty {
			r--
			continue
		}
		if !disk[l].empty {
			l++
			continue
		}
		disk[l] = disk[r]
		disk[r] = cell{empty: true}
		l++
		r--
	}
	return checksum(disk)
}

func part2(disk []cell) (res int) {
	empties := findEmpty(disk)

	r := len(disk) - 1
	for r >= 0 {
		if disk[r].empty {
			r--
			continue
		}
		if disk[r].moved {
			r--
			continue
		}
		j := r - 1
		for ; j >= 0; j-- {
			if disk[j].empty || disk[j].digit != disk[r].digit {
				break
			}
		}
		l := r - j
		for p, e := range empties {
			if e.index > r {
				break
			}
			if l > e.free {
				continue
			}
			for i := 0; i < l; i++ {
				disk[e.index+i] = disk[r-i]
				disk[e.index+i].moved = true
				disk[r-i] = cell{empty: true}
			}
			if l == e.free {
				empties = append(empties[:p], empties[p+1:]...)
			} else {
				empties[p].free = e.free - l
				empties[p].index += l
			}
			break
		}
		r = r - l
	}

	return checksum(disk)
}

func parseInput(input string) []cell {
	var disk []cell
	file := true
	fid := 0
	for i := 0; i < len(input); i++ {
		n := int(input[i] - '0')
		if file {
			for j := 0; j < n; j++ {
				disk = append(disk, cell{digit: fid})
			}
			fid++
		} else {
			for j := 0; j < n; j++ {
				disk = append(disk, cell{empty: true})
			}
		}
		file = !file
	}
	return disk
}

type cell struct {
	empty bool
	digit int
	moved bool
}

type empty struct {
	index int
	free  int
}

func checksum(d []cell) int {
	var sum int
	for pos, c := range d {
		if c.empty {
			continue
		}
		sum += pos * c.digit
	}
	return sum
}

func printDisk(d []cell) {
	for _, c := range d {
		if c.empty {
			fmt.Print(".")
		} else {
			fmt.Print(c.digit)
		}
	}
	fmt.Println()
}

func findEmpty(disk []cell) []empty {
	var empties []empty
	for i := 0; i < len(disk); i++ {
		if !disk[i].empty {
			continue
		}
		j := i + 1
		for ; j < len(disk); j++ {
			if !disk[j].empty {
				break
			}
		}
		empties = append(empties, empty{index: i, free: j - i})
		i = j
	}
	return empties
}
