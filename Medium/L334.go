package Medium

func increasingTriplet(nums []int) bool {
	left, mid := math.MaxInt32, math.MaxInt32

	for _, n := range nums {
		if n > mid {
			return true
		}

		if n < mid && n > left {
			mid = n
		} else if n < left {
			left = n
		}
	}

	return false
}
