package Medium

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for i:= 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}

	return count(&dp, 0, 0, m, n)
}

func isValid(i, j, m, n int) bool {
	return i >= 0 && i < m && j >=0 && j < n
}

func count(dp *[][]int, i, j, m, n int) int {
	if !isValid(i, j, m, n) {
		return 0
	}

	if i == m-1 && j == n-1 {
		return 1
	}

	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}

	(*dp)[i][j] = count(dp, i+1, j, m, n) + count(dp, i, j+1, m, n)
	
	return (*dp)[i][j]
}
