package Medium

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxSumAfterOperation(nums []int) int {
	mp, finalRes := make(map[string]int), nums[0]

	solve(0, &finalRes, true, nums, mp)
	return finalRes
}

func solve(idx int, finalRes *int, useSquare bool, nums []int, mp map[string]int) int {
	if idx >= len(nums) {
		return 0
	}

	key := fmt.Sprintf("%d-%v", idx, useSquare)
	if v, ok := mp[key]; ok {
		return v
	}

	ans := 0
	if useSquare {
		ans = max(nums[idx]*nums[idx], (nums[idx]*nums[idx])+solve(idx+1, finalRes, false, nums, mp))
	}

	ans = max(ans, max(nums[idx]+solve(idx+1, finalRes, useSquare, nums, mp), nums[idx]))

	mp[key] = ans
	*finalRes = max(*finalRes, ans)

	return mp[key]
}
