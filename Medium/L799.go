package Medium

func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([][]float64, query_row+1)
	for i := 0; i < query_row+1; i++ {
		dp[i] = make([]float64, query_glass+1)
		for j := 0; j < query_glass+1; j++ {
			dp[i][j] = -1
		}
	}

	ans := solve(query_row, query_glass, poured, &dp)

	if ans > 1 {
		return 1
	}

	return ans
}

func solve(row, col, poured int, dp *[][]float64) float64 {
	if col < 0 || col > row {
		return 0
	}
	if row == 0 && col == 0 {
		return float64(poured)
	}

	if (*dp)[row][col] != -1 {
		return (*dp)[row][col]
	}

	(*dp)[row][col] = max(solve(row-1, col-1, poured, dp)-1, 0)/2 + max(solve(row-1, col, poured, dp)-1, 0)/2

	return (*dp)[row][col]
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}
