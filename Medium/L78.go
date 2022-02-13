package Medium

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}

	bt(&ans, make([]int, 0), nums, 0)

	return ans
}

func bt(ans *[][]int, tmp, nums []int, start int) {
	copyTmp := make([]int, len(tmp))
	copy(copyTmp, tmp)
	*ans = append(*ans, copyTmp)

	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		bt(ans, tmp, nums, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}
