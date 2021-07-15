package Medium

func triangleNumber(nums []int) int {
	res := 0
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 2; i < len(nums); i++ {
		left, right := 0, i-1
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				res += right - left
				right--
			} else {
				left++
			}
		}
	}

	return res
}
