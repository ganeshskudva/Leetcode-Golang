package Medium

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int

	if len(candidates) == 0 {
		return ans
	}

	bt(&ans, make([]int, 0), candidates, 0, target)

	return ans
}

func bt(ans *[][]int, tmp, nums []int, idx, tgt int) {
	if tgt < 0 || idx > len(nums) {
		return
	}

	if tgt == 0 {
		cpyTmp := make([]int, len(tmp))
		copy(cpyTmp, tmp)
		*ans = append(*ans, cpyTmp)
	}

	for i := idx; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		bt(ans, tmp, nums, i, tgt-nums[i])
		tmp = tmp[:len(tmp)-1]
	}
}
