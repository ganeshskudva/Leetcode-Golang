package Medium

func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	var tmp []int

	if len(nums) == 0 {
		return ans
	}

	sort.Ints(nums)
	bt(&ans, tmp, nums, 0)

	return ans
}

func bt(ans *[][]int, tmp []int, nums []int, start int) {
	newRow := make([]int, len(tmp))
	copy(newRow, tmp)
	*ans = append(*ans, newRow)
	for i := start; i < len(nums); i++ {
		if i > 0 && i > start && nums[i] == nums[i-1] {
			continue
		}
		tmp = append(tmp, nums[i])
		bt(ans, tmp, nums, i+1)
		tmp = (tmp)[:len(tmp)-1]
	}
}
