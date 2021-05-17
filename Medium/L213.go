package Medium

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robHelper(nums[:len(nums)-1]), robHelper(nums[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func robHelper(nums []int) int {
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
