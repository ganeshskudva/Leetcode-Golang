package Medium

func getMaximumGold(grid [][]int) int {
	var (
		isValid func(x, y int) bool
		solve   func(x, y int) int
	)
	m, n, res := len(grid), len(grid[0]), 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	isValid = func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] != 0
	}

	solve = func(x, y int) int {
		if !isValid(x, y) {
			return 0
		}

		curr := grid[x][y]
		grid[x][y] = 0
		maxGold := 0
		for _, d := range dirs {
			maxGold = max(maxGold, solve(x+d[0], y+d[1]))
		}

		grid[x][y] = curr
		return curr + maxGold
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				res = max(res, solve(i, j))
			}
		}
	}

	return res
}
