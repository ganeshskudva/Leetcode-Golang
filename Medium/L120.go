package Medium

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, len(triangle)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return count(&dp, triangle, 0, 0)

}

func isValid(triangle [][]int, i, j int) bool {
	m := len(triangle)
	if i < 0 || i >= m {
		return false
	}

	n := len(triangle[i])
	if j < 0 || j >= n {
		return false
	}

	return true
}

func count(dp *[][]int, triangle [][]int, i, j int) int {
	if !isValid(triangle, i, j) {
		return 0
	}
	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}

	(*dp)[i][j] = triangle[i][j] + min(count(dp, triangle, i+1, j), count(dp, triangle, i+1, j+1))

	return (*dp)[i][j]
}
