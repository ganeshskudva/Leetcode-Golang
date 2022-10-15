package Hard

func getLengthOfOptimalCompression(s string, k int) int {
	m, n := len(s), k+1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	return solve(s, 0, k, &dp)
}

func solve(s string, currIdx, rest int, dp *[][]int) int {
	if currIdx == len(s) || len(s)-currIdx <= rest {
		return 0
	}

	if (*dp)[currIdx][rest] != -1 {
		return (*dp)[currIdx][rest]
	}

	free := make([]int, 26)
	most, res := 0, math.MaxInt32

	for i := currIdx; i < len(s); i++ {
		idx := s[i] - 'a'
		free[idx]++

		if most < free[idx] {
			most = free[idx]
		}
		if rest >= i-currIdx+1-most {
			res = min(res, getLen(most)+1+solve(s, i+1, rest-(i-currIdx+1-most), dp))
		}
	}

	(*dp)[currIdx][rest] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getLen(most int) int {
	if most == 1 {
		return 0
	}

	if most < 10 {
		return 1
	}

	if most < 99 {
		return 2
	}

	return 3
}
