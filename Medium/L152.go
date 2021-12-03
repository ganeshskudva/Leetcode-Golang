package Medium

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxi, mini, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := maxi
		maxi = max(max(maxi*nums[i], mini*nums[i]), nums[i])
		mini = min(min(tmp*nums[i], mini*nums[i]), nums[i])

		res = max(maxi, res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
