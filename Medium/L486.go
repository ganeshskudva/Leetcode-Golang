package Medium

func PredictTheWinner(nums []int) bool {
	var solve func(i, j, c int) int
	mp := make(map[string]int)

	solve = func(i, j, c int) int {
		if j < i {
			return 0
		}

		key := fmt.Sprintf("%d-%d-%d", i, j, c)
		if v, ok := mp[key]; ok {
			return v
		}

		res := 0
		if c == 0 {
			res = max(solve(i+1, j, 1-c)+nums[i], solve(i, j-1, 1-c)+nums[j])
		} else {
			res = min(solve(i+1, j, 1-c)-nums[i], solve(i, j-1, 1-c)-nums[j])
		}

		return res
	}

	ans := solve(0, len(nums)-1, 0)
	return ans >= 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
