package Easy

func findMaxConsecutiveOnes(nums []int) int {
	result, max := 0, 0

	if len(nums) == 0 {
		return result
	}

	for i := range nums {
		max += nums[i]
		if nums[i] == 0 {
			max = 0
		} else {
			if max >= result {
				result = max
			}
		}
	}
	return result
}
