package Easy

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	var res []int
	for i := range nums {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}
