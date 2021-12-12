package Medium

func canPartition(nums []int) bool {
	sum, n := 0, len(nums)

	for _, itm := range nums {
		sum += itm
	}

	if sum%2 != 0 {
		return false
	}

	sum /= 2
	dp, visited := make([][]bool, sum+1), make([][]bool, sum+1)
	for i := 0; i < sum+1; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < sum+1; i++ {
		visited[i] = make([]bool, n)
	}
	return solve(nums, dp, visited, sum, n-1)
}

func solve(nums []int, dp, visited [][]bool, sum, index int) bool {
	if index < 0 || sum < 0 {
		return false
	}
	if sum == 0 {
		dp[sum][index] = true
		return dp[sum][index]
	}
	if visited[sum][index] {
		return dp[sum][index]
	}
	visited[sum][index] = true
	dp[sum][index] = solve(nums, dp, visited, sum, index-1) || solve(nums, dp, visited, sum-nums[index], index-1)
	return dp[sum][index]
}
