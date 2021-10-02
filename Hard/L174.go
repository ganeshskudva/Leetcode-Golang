package Hard

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt64
		}
	}

	return solve(dungeon, 0, 0, &dp)
}

func solve(mat [][]int, i, j int, dp *[][]int) int {
	m, n := len(mat), len(mat[0])

	if i == m || j == n {
		return math.MaxInt64
	}
	if i == m-1 && j == n-1 {
		if mat[i][j] < 0 {
			return -mat[i][j] + 1
		} else {
			return 1
		}
	}

	if (*dp)[i][j] != math.MaxInt64 {
		return (*dp)[i][j]
	}

	right, down := solve(mat, i, j+1, dp), solve(mat, i+1, j, dp)
	minHpReq := min(right, down) - mat[i][j]

	(*dp)[i][j] = 1
	if minHpReq > 0 {
		(*dp)[i][j] = minHpReq
	}

	return (*dp)[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
