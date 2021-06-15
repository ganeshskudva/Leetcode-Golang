package Medium

func makesquare(matchsticks []int) bool {
	sum, max, k, vis := 0, 0, 4, make([]bool, len(matchsticks))
	for _, n := range matchsticks {
		sum += n
		if n > max {
			max = n
		}
	}

	if sum%k != 0 || max > sum/k {
		return false
	}

	return solve(matchsticks, &vis, k, sum/k, 0, 0)
}

func solve(nums []int, vis *[]bool, k, targetSubsetSum, curSubsetSum, nextIndexToCheck int) bool {
	if k == 0 {
		return true
	}

	if curSubsetSum == targetSubsetSum {
		return solve(nums, vis, k-1, targetSubsetSum, 0, 0)
	}

	for i := nextIndexToCheck; i < len(nums); i++ {
		if !(*vis)[i] && curSubsetSum+nums[i] <= targetSubsetSum {
			(*vis)[i] = true
			if solve(nums, vis, k, targetSubsetSum, curSubsetSum+nums[i], i+1) {
				return true
			}
			(*vis)[i] = false
		}
	}

	return false
}
