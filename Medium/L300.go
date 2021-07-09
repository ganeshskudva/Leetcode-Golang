package Medium

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = math.MinInt64
	}

	return solve(nums, 0, math.MinInt64, &dp)
}

func solve(nums []int, idx, prev int, dp *[]int) int {
	if idx >= len(nums) {
		return 0
	}

	curr := nums[idx]
	left, right := 0, 0

	if curr > prev {
		if (*dp)[idx] == math.MinInt64 {
			(*dp)[idx] = 1 + solve(nums, idx+1, curr, dp)
		}
		left = (*dp)[idx]
	}

	right = solve(nums, idx+1, prev, dp)

	if left > right {
		return left
	}

	return right
}
