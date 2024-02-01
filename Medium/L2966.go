package Medium

func divideArray(nums []int, k int) [][]int {
	var res [][]int
	if len(nums)%3 != 0 {
		return res
	}

	sort.Ints(nums)
	groupIdx := 0
	for i := 0; i < len(nums); i += 3 {
		if i+2 < len(nums) && nums[i+2]-nums[i] <= k {
			res = append(res, []int{nums[i], nums[i+1], nums[i+2]})
			groupIdx++
		} else {
			return [][]int{}
		}
	}

	return res
}
