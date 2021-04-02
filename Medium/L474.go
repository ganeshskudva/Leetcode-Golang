package Medium

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		one, zero := 0, 0
		for _, c := range s {
			if c == '1' {
				one++
			} else {
				zero++
			}
		}
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				if one <= j && zero <= i {
					dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
				}
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
