package Medium

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if k <= 0 || sum%k != 0 {
		return false
	}
	visited := make([]bool, len(nums))
	return solve(nums, &visited, 0, k, 0, 0, sum/k)
}

func solve(nums []int, visited *[]bool, startIdx, k, curSum, curNum, target int) bool {
	if k == 1 {
		return true
	}

	if curSum == target && curNum > 0 {
		return solve(nums, visited, 0, k-1, 0, 0, target)
	}

	for i := startIdx; i < len(nums); i++ {
		if !(*visited)[i] {
			(*visited)[i] = true
			if solve(nums, visited, i+1, k, curSum+nums[i], curNum+1, target) {
				return true
			}
			(*visited)[i] = false
		}
	}

	return false
}
