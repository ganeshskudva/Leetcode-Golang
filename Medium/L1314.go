package Medium

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n, total := len(mat), len(mat[0]), make([][]int, len(mat)+1)
	for i := 0; i <= m; i++ {
		total[i] = make([]int, n+1)
	}

	for r := 1; r <= m; r++ {
		for c := 1; c <= n; c++ {
			total[r][c] = mat[r-1][c-1] + total[r-1][c] + total[r][c-1] - total[r-1][c-1]
		}
	}

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			r1, c1 := max(0, r-k), max(0, c-k)
			r2, c2 := min(m-1, r+k), min(n-1, c+k)

			r1, c1 = r1+1, c1+1
			r2, c2 = r2+1, c2+1
			res[r][c] = total[r2][c2] - total[r2][c1-1] - total[r1-1][c2] + total[r1-1][c1-1]
		}
	}

	return res
}
