package Medium

func onesMinusZeros(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	diff, rows, cols := make([][]int, m), make([]int, m), make([]int, n)

	for i := 0; i < m; i++ {
		diff[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rows[i] += grid[i][j]
			cols[j] += grid[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff[i][j] = rows[i] + cols[j] - (n - rows[i]) - (m - cols[j])
		}
	}
	
	return diff
}
