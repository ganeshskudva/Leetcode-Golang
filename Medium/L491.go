package Medium

func findSubsequences(nums []int) [][]int {
	var (
		tmp []int
		res [][]int
	)
	solve(0, nums, &tmp, &res)
	return res
}

func solve(idx int, nums []int, tmp *[]int, res *[][]int) {
	if idx > len(nums)-1 {
		if len(*tmp) > 1 {
			tmpCopy := make([]int, len(*tmp))
			copy(tmpCopy, *tmp)
			*res = append(*res, tmpCopy)
		}
		return
	}

	if len(*tmp) == 0 || nums[idx] >= (*tmp)[len(*tmp)-1] {
		*tmp = append(*tmp, nums[idx])
		solve(idx+1, nums, tmp, res)
		*tmp = (*tmp)[:len(*tmp)-1]
	}

	if idx > 0 && len(*tmp) > 0 && nums[idx] == (*tmp)[len(*tmp)-1] {
		return
	}
	solve(idx+1, nums, tmp, res)
}
