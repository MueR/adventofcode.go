package util

import (
	"github.com/MueR/adventofcode.go/maths"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Manhattan(q Point) int {
	return maths.AbsInt(p.X-q.X) + maths.AbsInt(p.Y-q.Y)
}
