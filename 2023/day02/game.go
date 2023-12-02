package main

import (
	"strconv"
	"strings"

	"github.com/MueR/adventofcode.go/maths"
)

type game struct {
	num   int
	hands []hand
	min   hand
}

type hand struct {
	red   int
	green int
	blue  int
}

func newGame(input string) (g game) {
	gs := strings.Split(input, ":")
	g.num, _ = strconv.Atoi(gs[0][5:])
	for _, set := range strings.Split(gs[1], ";") {
		for _, cubes := range strings.Split(set, ",") {
			d := strings.TrimSpace(cubes)
			cube := strings.Split(d, " ")
			h := hand{}
			num, _ := strconv.Atoi(cube[0])
			switch cube[1] {
			case "red":
				h.red = num
				g.min.red = maths.MaxInt([]int{g.min.red, num})
			case "green":
				h.green = num
				g.min.green = maths.MaxInt([]int{g.min.green, num})
			case "blue":
				h.blue = num
				g.min.blue = maths.MaxInt([]int{g.min.blue, num})
			}
			g.hands = append(g.hands, h)
		}
	}
	return g
}

func (h hand) Pow() int {
	return h.red * h.green * h.blue
}
func (h hand) Valid(rules hand) bool {
	return h.red <= rules.red && h.green <= rules.green && h.blue <= rules.blue
}
