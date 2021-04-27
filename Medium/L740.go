package Medium

func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
	dp := make([]int, len(nums)+1)
	for i := range dp {
		dp[i] = -1
	}
	
	return solve(nums, 0, &dp)
}

func solve(nums []int, idx int, dp *[]int) int {
	if idx == len(nums) {
		return 0
	}
	
	if (*dp)[idx] != -1 {
		return (*dp)[idx]
	}
	
	elem, sum, next := nums[idx], 0, idx 
	
	for next < len(nums) && nums[next] == elem {
		next++
		sum += elem
	}
	
	for next < len(nums) && nums[next] == elem+1 {
		next++
	}
	
	(*dp)[idx] = max(sum + solve(nums, next, dp), solve(nums, idx+1, dp))
	
	return (*dp)[idx]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
