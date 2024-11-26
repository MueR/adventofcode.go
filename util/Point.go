package util

import (
	"math"

	"github.com/MueR/adventofcode.go/maths"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Manhattan(q Point) int {
	return maths.AbsInt(p.X-q.X) + maths.AbsInt(p.Y-q.Y)
}
func (p *Point) Pythagorean(q Point) int {
	return int(math.Sqrt(math.Pow(float64(p.X-q.X), 2) + math.Pow(float64(p.Y-q.Y), 2)))
}
