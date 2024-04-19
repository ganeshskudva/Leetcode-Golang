package Medium

func numIslands(grid [][]byte) int {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	cnt, m, n, mp := 0, len(grid), len(grid[0]), make(map[string]bool)
	var (
		getKey  func(i, j int) string
		isValid func(x, y int) bool
		solve   func(i, j int)
	)

	getKey = func(i, j int) string {
		return fmt.Sprintf("%d,%d", i, j)
	}

	isValid = func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1'
	}

	solve = func(i, j int) {
		if !isValid(i, j) {
			return
		}

		key := getKey(i, j)
		if mp[key] {
			return
		}
		mp[key] = true

		for _, d := range dirs {
			solve(i+d[0], j+d[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				if _, ok := mp[getKey(i, j)]; !ok {
					cnt++
					solve(i, j)
				}
			}
		}
	}

	return cnt
}
