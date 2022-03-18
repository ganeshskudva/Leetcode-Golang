package Easy

func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	if len(nums) == 0 {
		return ret
	}

	i, j := 0, len(nums)-1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		if abs(nums[i]) > abs(nums[j]) {
			ret[idx] = nums[i] * nums[i]
			i++
		} else {
			ret[idx] = nums[j] * nums[j]
			j--
		}
	}

	return ret
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
