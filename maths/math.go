package maths

import (
	"math"
)

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
