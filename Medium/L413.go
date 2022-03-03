package Medium

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	
	cnt := 0
	for i := 0; i < len(nums)-2; i++ {
		d := nums[i+1] - nums[i]
		cnt += helper(nums, i+2, nums[i+1], d, len(nums))
	}
	return cnt
}

func helper(nums []int, idx, prev, d, n int) int {
	if idx == n || nums[idx] - prev != d {
		return 0
	}
	
	return 1 + helper(nums, idx+1, nums[idx], d, n)
}
