package math

func SumIntSlice(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func MaxInt(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
