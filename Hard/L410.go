package Hard

func splitArray(nums []int, m int) int {
	if len(nums) == 0 || m == 0 {
		return 0
	}

	max, sum := math.MinInt32, 0
	for _, n := range nums {
		if n > max {
			max = n
		}
		sum += n
	}

	l, r, mid := max, sum, 0

	for l < r {
		mid = l + (r-l)/2
		if isValid(nums, mid, m) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func isValid(nums []int, cap, m int) bool {
	days, sum := 1, 0

	for _, n := range nums {
		sum += n
		if sum > cap {
			sum = n
			days++
		}
	}

	return days <= m
}
