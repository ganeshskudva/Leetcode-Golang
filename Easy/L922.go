package Easy

func sortArrayByParityII(nums []int) []int {
	i, j, n := 0, 1, len(nums)

	for i < n && j < n {
		for i < n && nums[i]%2 == 0 {
			i += 2
		}
		for j < n && nums[j]%2 == 1 {
			j += 2
		}
		if i < n && j < n {
			// swap
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return nums
}
