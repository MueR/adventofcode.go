package grid_test

import (
	"strings"
	"testing"

	"github.com/MueR/adventofcode.go/data-structures/grid"
)

var example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestStringGrid(t *testing.T) {
	g := grid.NewGrid[string](strings.Split(example, "\n"))
	_ = g
}
