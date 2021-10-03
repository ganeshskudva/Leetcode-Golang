package Medium

func canJump(nums []int) bool {
	n, dp := len(nums), make([]int, len(nums))
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	return solve(nums, 0, &dp)
}

func solve(nums []int, idx int, dp *[]int) bool {
	if idx >= len(nums)-1 {
		return true
	}

	if nums[idx] == 0 {
		(*dp)[idx] = 0
		return false
	}

	if (*dp)[idx] != -1 {
		return (*dp)[idx] == 1
	}

	jumps := nums[idx]
	for i := 1; i <= jumps; i++ {
		if solve(nums, idx+i, dp) {
			(*dp)[idx] = 1
			return true
		}
	}

	(*dp)[idx] = 0
	return false
}
