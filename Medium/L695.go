package Medium

func maxAreaOfIsland(grid [][]int) int {
	max := math.MinInt32
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	cnt := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				cnt = 0
				dfs(grid, &vis, &cnt, i, j, dirs)
			}
			if max < cnt {
				max = cnt
			}
		}
	}

	if max == math.MinInt32 {
		return 0
	}

	return max
}

func isValid(grid [][]int, x, y int) bool {
	m, n := len(grid), len(grid[0])

	return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1
}

func dfs(grid [][]int, vis *[][]bool, cnt *int, i, j int, dirs [][]int) {
	if !isValid(grid, i, j) {
		return
	}

	if (*vis)[i][j] {
		return
	}

	(*vis)[i][j] = true
	*cnt++

	for _, d := range dirs {
		dfs(grid, vis, cnt, i+d[0], j+d[1], dirs)
	}

}
