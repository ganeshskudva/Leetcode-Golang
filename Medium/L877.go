package Medium

func stoneGame(piles []int) bool {
	m, n := len(piles), len(piles)
	memo := make([][]int, m*n)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	sum := 0
	for _, v := range piles {
		sum += v
	}

	return solve(piles, &memo, 0, len(piles)-1) > sum/2
}

func solve(piles []int, memo *[][]int, startIdx, endIdx int) int {
	if startIdx > endIdx {
		return 0
	}

	if (*memo)[startIdx][endIdx] > 0 {
		return (*memo)[startIdx][endIdx]
	}

	if startIdx+1 == endIdx {
		(*memo)[startIdx][endIdx] = max(piles[startIdx], piles[endIdx])
	} else {
		a := piles[startIdx] + min(solve(piles, memo, startIdx+2, endIdx), solve(piles, memo, startIdx+1, endIdx-1))
		b := piles[endIdx] + min(solve(piles, memo, startIdx+1, endIdx-1), solve(piles, memo, startIdx, endIdx-2))
		(*memo)[startIdx][endIdx] = max(a, b)
	}

	return (*memo)[startIdx][endIdx]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
