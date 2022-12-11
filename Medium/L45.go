package Medium

func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}

	return solve(nums, &dp, 0)
}

func solve(nums []int, dp *[]int, idx int) int {
	if idx >= len(nums)-1 {
		return 0
	}

	if (*dp)[idx] != -1 {
		return (*dp)[idx]
	}

	minJump := len(nums)
	for i := idx + 1; i <= idx+nums[idx]; i++ {
		minJump = min(minJump, 1+solve(nums, dp, i))
	}

	(*dp)[idx] = minJump
	return minJump
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
