package Hard

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
			for k := 0; k < n; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	return dfs(grid, m, n, 0, 0, n-1, &dp)
}

func dfs(grid [][]int, m, n, r, c1, c2 int, dp *[][][]int) int {
	if r == m {
		return 0
	}

	if (*dp)[r][c1][c2] != -1 {
		return (*dp)[r][c1][c2]
	}

	ans := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			nc1, nc2 := c1+i, c2+j
			if nc1 >= 0 && nc1 < n && nc2 >= 0 && nc2 < n {
				val := dfs(grid, m, n, r+1, nc1, nc2, dp)
				if val > ans {
					ans = val
				}
			}
		}
	}

	cherries := 0
	if c1 == c2 {
		cherries = grid[r][c1]
	} else {
		cherries = grid[r][c1] + grid[r][c2]
	}

	(*dp)[r][c1][c2] = ans + cherries
	return (*dp)[r][c1][c2]
}
