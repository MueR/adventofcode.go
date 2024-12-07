package maths

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func SumIntSlice(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return sum
}

func MultiplyIntSlice(nums []int) (product int) {
	product = 1
	for _, n := range nums {
		product *= n
	}
	return product
}

func MaxInt(nums []int) (max int) {
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func MinInt(nums ...int) (min int) {
	min = math.MaxInt64
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

func AbsInt(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

// GCD finds greatest common divisor via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple  via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Sign[T Number](v T) int {
	switch {
	case v < 0:
		return -1
	case v == 0:
		return 0
	default:
		return 1
	}
}

func Concat[T constraints.Integer](a, b T) T {
	pow := T(10)
	for b >= pow {
		pow *= 10
	}
	return a*pow + b
}
