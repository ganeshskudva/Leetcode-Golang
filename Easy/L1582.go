package Easy

func numSpecial(mat [][]int) int {
	m, n, res, col, row := len(mat), len(mat[0]), 0, make([]int, len(mat[0])), make([]int, len(mat))

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				col[j], row[i] = col[j]+1, row[i]+1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && row[i] == 1 && col[j] == 1 {
				res++
			}
		}
	}

	return res
}
