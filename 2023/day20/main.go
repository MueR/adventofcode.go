package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/maths"
)

var (
	//go:embed input.txt
	input string
)

const (
	FlipFlop ModuleType = iota
	Conjunction
	Broadcaster
)

type (
	ModuleType int
	Module     struct {
		Name         string
		Destinations []string
		Type         ModuleType
		Conjunction  []Pulse
		State        bool
	}
	Pulse struct {
		from, to string
		isHigh   bool
	}
)

func (m *Module) ReceivePulse(in Pulse) (out []Pulse) {
	switch m.Type {
	case FlipFlop:
		if in.isHigh {
			return []Pulse{}
		} else {
			m.State = !m.State
			in.isHigh = m.State
		}
	case Conjunction:
		for i := range m.Conjunction {
			if m.Conjunction[i].from == in.from {
				m.Conjunction[i].isHigh = in.isHigh
				break
			}
		}
		low := false
		for i := range m.Conjunction {
			if !m.Conjunction[i].isHigh {
				low = true
				break
			}
		}
		in.isHigh = low
	default:
		// Broadcaster
	}
	out = make([]Pulse, len(m.Destinations))
	for i := range m.Destinations {
		out[i] = Pulse{
			from:   m.Name,
			to:     m.Destinations[i],
			isHigh: in.isHigh,
		}
	}
	return out
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
	modules, _ := parseInput(input)
	var low, high int
	for i := 0; i < 1000; i++ {
		pulses := []Pulse{
			{from: "btn", to: "broadcaster"},
		}
		for len(pulses) > 0 {
			for _, p := range pulses {
				if p.isHigh {
					high++
				} else {
					low++
				}
			}
			var newPulses []Pulse
			for _, p := range pulses {
				for current := range modules {
					if modules[current].Name == p.to {
						newPulses = append(newPulses, modules[current].ReceivePulse(p)...)
					}
				}
			}
			pulses = newPulses
		}
	}
	return low * high
}

func part2(input string) int {
	modules, finalMods := parseInput(input)
	i := 0
	multiples := make([][]int, len(finalMods))
	for {
		pulses := []Pulse{
			{from: "btn", to: "broadcaster"},
		}
		for len(pulses) > 0 {
			var newPulses []Pulse
			for _, p := range pulses {
				for current := range modules {
					if modules[current].Name == p.to {
						newPulses = append(newPulses, modules[current].ReceivePulse(p)...)
					}
				}
			}
			for _, v := range newPulses {
				if !v.isHigh {
					continue
				}
				idx := slices.Index(finalMods, v.from)
				if idx == -1 {
					continue
				}
				multiples[idx] = append(multiples[idx], i)
				completeSet := true
				for _, m := range multiples {
					if len(m) < 3 {
						completeSet = false
					}
				}
				if completeSet {
					f := make([]int, len(finalMods))
					for i := range multiples {
						f[i] = multiples[i][1] - multiples[i][0]
					}
					return maths.LCM(f[0], f[1], f[2:]...)
				}
			}
			pulses = newPulses
		}
		i++
	}
}

func parseInput(input string) (mods []Module, finalModules []string) {
	var finalMod Module
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		mt := Broadcaster
		name := parts[0]
		switch name[0] {
		case '&':
			mt = Conjunction
			name = name[1:]
		case '%':
			mt = FlipFlop
			name = name[1:]
		}
		mod := Module{
			Name:         name,
			Destinations: strings.Split(parts[1], ", "),
			Type:         mt,
			State:        false,
		}
		if slices.Contains(mod.Destinations, "rx") {
			finalMod = mod
		}
		mods = append(mods, mod)
	}
	for cur := range mods {
		if slices.Contains(mods[cur].Destinations, finalMod.Name) {
			finalModules = append(finalModules, mods[cur].Name)
		}
		if mods[cur].Type != Conjunction {
			continue
		}
		for other := range mods {
			for _, dest := range mods[other].Destinations {
				if dest == mods[cur].Name {
					mods[cur].Conjunction = append(mods[cur].Conjunction, Pulse{from: mods[other].Name})
				}
			}
		}
	}
	return mods, finalModules
}
