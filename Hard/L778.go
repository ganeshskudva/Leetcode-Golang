package Hard

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func swimInWater(grid [][]int) int {
	time, m := 0, len(grid)

	for {
		vis := make(map[int]bool)
		dfs(grid, vis, 0, 0, time)
		time++
		if _, ok := vis[m*m-1]; ok {
			break
		}
	}

	return time - 1
}

func isValid(grid [][]int, x, y, time int) bool {
	m, n := len(grid), len(grid[0])

	return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] <= time
}

func dfs(grid [][]int, vis map[int]bool, i, j, time int) {
	if !isValid(grid, i, j, time) {
		return
	}

	if _, ok := vis[i*len(grid)+j]; ok {
		return
	}

	vis[i*len(grid)+j] = true

	for _, d := range dirs {
		dfs(grid, vis, i+d[0], j+d[1], time)
	}
}
