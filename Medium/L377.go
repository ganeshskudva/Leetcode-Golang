package Medium

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	for i, _ := range dp {
		dp[i] = -1
	}
	dp[0] = 1
	return bt(&dp, nums, target)
}

func bt(dp *[]int, nums []int, tgt int) int {
	if (*dp)[tgt] != -1 {
		return (*dp)[tgt]
	}

	res := 0
	for _, n := range nums {
		if tgt >= n {
			res += bt(dp, nums, tgt-n)
		}
	}

	(*dp)[tgt] = res
	return (*dp)[tgt]
}
