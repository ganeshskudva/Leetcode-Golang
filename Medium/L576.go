package Medium

const MOD = 1000000007

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, maxMove+1)
			for k := 0; k < maxMove+1; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	return dfs(&dp, m, n, startRow, startColumn, maxMove)
}

func isValid(m, n, x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func dfs(dp *[][][]int, m, n, i, j, maxMove int) int {
	if maxMove < 0 {
		return 0
	}

	if !isValid(m, n, i, j) {
		return 1
	}

	if (*dp)[i][j][maxMove] != -1 {
		return (*dp)[i][j][maxMove]
	}

	sum := 0
	for _, d := range dirs {
		sum += dfs(dp, m, n, i+d[0], j+d[1], maxMove-1) % MOD
	}

	(*dp)[i][j][maxMove] = sum % MOD

	return (*dp)[i][j][maxMove] % MOD
}
