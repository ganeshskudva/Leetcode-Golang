package Medium

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return solve(word1, word2, 0, 0, &dp)
}

func solve(w1, w2 string, p1, p2 int, dp *[][]int) int {
	m, n := len(w1), len(w2)

	if p1 == m {
		return n - p2
	}
	if p2 == n {
		return m - p1
	}

	if (*dp)[p1][p2] != -1 {
		return (*dp)[p1][p2]
	}

	if w1[p1] == w2[p2] {
		(*dp)[p1][p2] = solve(w1, w2, p1+1, p2+1, dp)
	} else {
		(*dp)[p1][p2] = 1 + min(solve(w1, w2, p1+1, p2, dp), solve(w1, w2, p1, p2+1, dp))
	}
	
	return (*dp)[p1][p2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	
	return b
}
