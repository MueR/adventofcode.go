package tilemap

import (
	"fmt"
	"io"
	"iter"
	"slices"

	"github.com/MueR/adventofcode.go/cast"
	"github.com/MueR/adventofcode.go/maths"
	"github.com/MueR/adventofcode.go/util"
	"github.com/beefsack/go-astar"
)

type (
	Container[T comparable] struct {
		Value    T
		tileMap  *Map[T]
		position util.Point
	}

	Map[T comparable] struct {
		tiles []Container[T]
		w     int
		h     int

		CostFunc      func(a, b Container[T]) float64
		EstimateFunc  func(a, b Container[T]) float64
		NeighbourFunc func(container Container[T]) []Container[T]
	}
)

func (c Container[T]) Location() (int, int) {
	return c.position.X, c.position.Y
}

func (c Container[T]) PathNeighbors() (results []astar.Pather) {
	if c.tileMap.NeighbourFunc != nil {
		neighbours := c.tileMap.NeighbourFunc(c)
		results = make([]astar.Pather, len(neighbours))
		for i, v := range c.tileMap.NeighbourFunc(c) {
			results[i] = v
		}
	}
	return
}

func (c Container[T]) PathNeighborCost(to astar.Pather) float64 {
	if c.tileMap.CostFunc != nil {
		return c.tileMap.CostFunc(c, to.(Container[T]))
	}

	// Assume all paths are equal cost
	return 1
}

func (c Container[T]) PathEstimatedCost(to astar.Pather) float64 {
	toSpot := to.(Container[T])

	if c.tileMap.EstimateFunc != nil {
		return c.tileMap.EstimateFunc(c, toSpot)
	}

	return float64(maths.ManhattanDistance(c.position.X, c.position.Y, toSpot.position.X, toSpot.position.Y))
}

func FromInput(input io.Reader) *Map[rune] {
	return ConvertInputOf[rune](input, cast.ToRunes)
}

func ConvertInputOf[T comparable](input io.Reader, convert func(rune) T) *Map[T] {
	lines := slices.Collect(util.Lines(input))

	m := Of[T](len(lines[0]), len(lines))

	for row, line := range lines {
		for col, tile := range line {
			m.SetTile(col, row, convert(tile))
		}
	}

	return m
}

func Of[T comparable](w, h int) *Map[T] {
	return &Map[T]{
		tiles: make([]Container[T], w*h),
		w:     w,
		h:     h,
	}
}

func (t *Map[T]) Size() (int, int) {
	return t.w, t.h
}

func (t *Map[T]) outOfBounds(x, y int) bool {
	return x < 0 || y < 0 || x >= t.w || y >= t.h
}

func (t *Map[T]) indexOf(x, y int) (int, bool) {
	return x + (t.w * y), !t.outOfBounds(x, y)
}

func (t *Map[T]) SetTile(x, y int, tile T) {
	idx, ok := t.indexOf(x, y)
	if !ok {
		panic(fmt.Errorf("out of bounds tile access: [%d, %d] is not within the %dx%d map", x, y, t.w, t.h))
	}

	t.tiles[idx] = Container[T]{tileMap: t, Value: tile, position: util.Point{X: x, Y: y}}
}

func (t *Map[T]) ContainerAt(x, y int) (Container[T], bool) {
	idx, ok := t.indexOf(x, y)
	if !ok {
		return Container[T]{}, false
	}

	return t.tiles[idx], true
}

func (t *Map[T]) TileAt(x, y int) (T, bool) {
	c, ok := t.ContainerAt(x, y)
	return c.Value, ok
}

func (t *Map[T]) FirstContainerWith(v T) (Container[T], bool) {
	for _, c := range t.tiles {
		if c.Value == v {
			return c, true
		}
	}

	return Container[T]{}, false
}

func (t *Map[T]) AllContainersWith(v T) (results []Container[T]) {
	for _, c := range t.tiles {
		if c.Value == v {
			results = append(results, c)
		}
	}

	return results
}

// Values returns an iter.Seq2 over all values in the map starting at 0,0 one row at a time. Each tile will be visited
// exactly once.
func (t *Map[T]) Values() iter.Seq2[T, util.Point] {
	return func(yield func(T, util.Point) bool) {
		for i, c := range t.tiles {
			if !yield(c.Value, util.Point{X: i % t.w, Y: i / t.w}) {
				return
			}
		}
	}
}

// CardinalNeighbors returns neighbours in the N, S, E, W directions, if they exist
func (t *Map[T]) CardinalNeighbors(x, y int) iter.Seq2[T, util.Point] {
	return func(yield func(T, util.Point) bool) {
		for _, d := range []struct {
			x, y int
		}{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		} {
			if r, ok := t.TileAt(x+d.x, y+d.y); ok {
				if !yield(r, util.Point{X: x + d.x, Y: y + d.y}) {
					return
				}
			}
		}
	}
}

// AllNeighbors returns cardinal and diagonal neighbors if they exist
func (t *Map[T]) AllNeighbors(x, y int) iter.Seq2[T, util.Point] {
	return func(yield func(T, util.Point) bool) {
		// Start by walking cardinals
		for v, pos := range t.CardinalNeighbors(x, y) {
			if !yield(v, pos) {
				return
			}
		}

		// Then Walk diagonals
		for _, d := range []struct {
			x, y int
		}{
			{-1, -1},
			{-1, 1},
			{1, -1},
			{1, 1},
		} {
			if r, ok := t.TileAt(x+d.x, y+d.y); ok {
				if !yield(r, util.Point{X: x + d.x, Y: y + d.y}) {
					return
				}
			}
		}
	}
}

// PathBetween uses the A-Star path finding algorithm to find the most efficient path between the two locations in the
// map. By default, the following constraints are used for finding the path:
//
// * Neighbors on the cardinal direction are reachable if they exist in the tile map
// * The cost to traverse any neighbor is always 1
// * The cost estimate between any two points in the map is the manhattan distance between them
//
// To override these, override Map.NeighborFunc, Map.CostFunc, and Map.EstimateFunc
func (t *Map[T]) PathBetween(startX, startY, endX, endY int) ([]astar.Pather, float64, bool) {
	start, ok := t.ContainerAt(startX, startY)
	if !ok {
		return nil, 0, false
	}

	end, ok := t.ContainerAt(endX, endY)
	if !ok {
		return nil, 0, false
	}

	return astar.Path(start, end)
}
