package Medium

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var res []bool
	for i := 0; i < len(l); i++ {
		res = append(res, isArithmetic(nums, l[i], r[i]))
	}

	return res
}

func isArithmetic(nums []int, l, r int) bool {
	maximum, minimum := math.MinInt32, math.MaxInt32

	for i := l; i <= r; i++ {
		maximum = max(nums[i], maximum)
		minimum = min(nums[i], minimum)
	}

	length := r - l + 1
	visited := make([]bool, length)
	if (maximum-minimum)%(length-1) != 0 {
		return false
	}

	diff := (maximum - minimum) / (length - 1)
	if diff == 0 {
		return true
	}

	for i := l; i <= r; i++ {
		val := nums[i]
		if (val-minimum)%diff != 0 {
			return false
		} else {
			pos := (val - minimum) / diff
			if visited[pos] {
				return false
			}
			visited[pos] = true
		}
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
