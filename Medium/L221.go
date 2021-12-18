package Medium

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n, ans := len(matrix), len(matrix[0]), math.MinInt64
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				ans = max(ans, dfs(matrix, i, j, dp))
			}
		}
	}

	return ans * ans
}

func isValid(matrix [][]byte, x, y int) bool {
	m, n := len(matrix), len(matrix[0])

	return x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] == '1'
}

func dfs(matrix [][]byte, i, j int, dp [][]int) int {
	if !isValid(matrix, i, j) {
		return 0
	}

	if dp[i][j] != 0 {
		return dp[i][j]
	}

	down, right, diag := dfs(matrix, i+1, j, dp), dfs(matrix, i, j+1, dp), dfs(matrix, i+1, j+1, dp)
	dp[i][j] = 1 + min(min(down, right), diag)
	return dp[i][j]
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
