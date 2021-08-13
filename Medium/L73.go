package Medium

func setZeroes(matrix [][]int) {
	fr, fc, m, n := false, false, len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					fr = true
				}
				if j == 0 {
					fc = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if fr {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if fc {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
