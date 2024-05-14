package Medium

func matrixScore(grid [][]int) int {
	m, n, res := len(grid), len(grid[0]), 0
	
	res += (1 << (n - 1)) * m
	
	for j := 1; j < n; j++ {
		same := 0
		for i := 0; i < m; i++ {
			if grid[i][0] == grid[i][j] {
				same++
			}
		}
		res += (1 << (n - 1 - j))*max(same, m - same)
	}
	
	return res
}
