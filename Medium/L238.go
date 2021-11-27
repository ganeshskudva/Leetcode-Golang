package Medium

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	ret := make([]int, len(nums))
	ret[0] = 1
	for i := 1; i < len(nums); i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}

	right := 1
	for i := len(nums) - 2; i >= 0; i-- {
		right *= nums[i+1]
		ret[i] = ret[i] * right
	}

	return ret
}
