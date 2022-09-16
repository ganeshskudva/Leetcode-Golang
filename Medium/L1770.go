package Medium

func maximumScore(nums []int, multipliers []int) int {
	m := len(multipliers)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, m)
	}

	solveMaxScore(nums, multipliers, 0, 0, &memo)

	return memo[0][0]
}

func solveMaxScore(nums, multipliers []int, l, i int, memo *[][]int) int {
	if i == len(multipliers) {
		return 0
	}

	if (*memo)[l][i] != 0 {
		return (*memo)[l][i]
	}

	pickLeft := solveMaxScore(nums, multipliers, l+1, i+1, memo) + (nums[l] * multipliers[i])
	pickRight := solveMaxScore(nums, multipliers, l, i+1, memo) + (nums[len(nums)-(i-l)-1] * multipliers[i])

	(*memo)[l][i] = max(pickLeft, pickRight)
	return (*memo)[l][i]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
