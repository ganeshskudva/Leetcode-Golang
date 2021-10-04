package Easy

func islandPerimeter(grid [][]int) int {
	m, n, cnt, dirs := len(grid), len(grid[0]), 0, [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				for _, d := range dirs {
					x, y := i+d[0], j+d[1]
					if isValid(grid, x, y) {
						cnt++
					}
				}
			}
		}
	}

	return cnt
}

func isValid(grid [][]int, x, y int) bool {
	m, n := len(grid), len(grid[0])

	return x < 0 || y < 0 || x == m || y == n || grid[x][y] == 0
}
