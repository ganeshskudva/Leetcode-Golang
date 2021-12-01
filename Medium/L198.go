package Medium

func max(a, b int) int {
	if a > b {
		return a
	}
	
	return b
}

func rob(nums []int) int {
	mp := make(map[int]int)
	n := len(nums)
	return max(count(mp, nums, n-1), count(mp, nums, n-2))
}

func count(mp map[int]int, nums []int, n int) int {
	if _, ok := mp[n]; ok {
		return mp[n]
	}
	
	if n < 0 {
		return 0
	}
	
	mp[n] = nums[n] + max(count(mp, nums, n-2), count(mp, nums, n-3))
	
	return mp[n]
}
