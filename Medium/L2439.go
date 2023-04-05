package Medium

func minimizeArrayValue(nums []int) int {
	low, high, res := 0, math.MaxInt32, 0

	isPossible := func(maximum int) bool {
		excess := 0
		for _, n := range nums {
			if n > maximum {
				candidate := n - maximum
				if candidate > excess {
					return false
				}
				excess -= candidate
			} else {
				excess += maximum - n
			}
		}

		return true
	}

	for low <= high {
		mid := low + (high-low)/2
		if isPossible(mid) {
			high, res = mid-1, mid
		} else {
			low = mid + 1
		}
	}

	return res
}
