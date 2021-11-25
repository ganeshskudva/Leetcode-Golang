package Easy

func maxSubArray(nums []int) int {
	max := math.MinInt64
	solve(nums, 0, &max)

	return max
}

func solve(nums []int, idx int, max *int) int {
	maxSoFar := 0
	if idx == len(nums)-1 {
		maxSoFar = nums[idx]
	} else {
		maxSoFar = maximum(nums[idx], nums[idx]+solve(nums, idx+1, max))
	}
	*max = maximum(*max, maxSoFar)

	return maxSoFar
}

func maximum(a, b int) int {
	if a > b {
		return a
	}

	return b
}
