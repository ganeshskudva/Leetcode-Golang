package Medium

func permuteUnique(nums []int) [][]int {
	var (
		ans [][]int
		tmp []int
	)
	if len(nums) == 0 {
		return ans
	}

	sort.Ints(nums)
	used := make([]bool, len(nums))
	backTrack(&ans, &tmp, nums, &used)

	return ans
}

func backTrack(ans *[][]int, tmp *[]int, nums []int, used *[]bool) {
	if len(*tmp) == len(nums) {
		tmpCopy := make([]int, len(*tmp))
		copy(tmpCopy, *tmp)
		*ans = append(*ans, tmpCopy)
		return
	}

	for i := range nums {
		if (*used)[i] || (i > 0 && nums[i] == nums[i-1] && !(*used)[i-1]) {
			continue
		}

		(*used)[i] = true
		*tmp = append(*tmp, nums[i])
		backTrack(ans, tmp, nums, used)
		(*used)[i] = false
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
