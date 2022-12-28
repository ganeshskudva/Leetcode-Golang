package Medium

func findMaxConsecutiveOnes(nums []int) int {
	mp := make(map[string]int)
	
	return solve(0, 1, 0, nums, mp)
}

func solve(idx, k, cnt int, nums []int, mp map[string]int) int {
	if idx >= len(nums) {
		return cnt
	}

	if nums[idx] == 0 && k == 0 {
		return cnt
	}

	key := fmt.Sprintf("%d-%d", idx, k)
	if v, ok := mp[key]; ok {
		return v
	}

	if nums[idx] == 0 {
		mp[key] = max(solve(idx+1, k-1, cnt+1, nums, mp), solve(idx+1, k, 0, nums, mp))
	} else {
		mp[key] = solve(idx + 1, k, cnt + 1, nums, mp)
	}
	
	return mp[key]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
