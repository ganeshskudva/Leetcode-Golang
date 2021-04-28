package Hard

func uniquePathsIII(grid [][]int) int {
	x, y, empty := 0, 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				empty++
			} else if grid[i][j] == 1 {
				x, y = i, j
			}
		}
	}
	
	return dfs(&grid, x, y, -1, empty)
}

func isValid(grid *[][]int, x, y int) bool {
	m, n := len(*grid), len((*grid)[0])
	return x >= 0 && x < m && y >= 0 && y < n && (*grid)[x][y] != -1
}

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func dfs(grid *[][]int, i, j, cnt, need int) int {
	if !isValid(grid, i, j) {
		return 0
	}

	if (*grid)[i][j] == 2 {
		if cnt == need {
			return 1
		} else {
			return 0
		}
	}
	(*grid)[i][j] = -1
	sum := 0
	for _, d := range dirs {
		sum += dfs(grid, i+d[0], j+d[1], cnt+1, need)
	}
	(*grid)[i][j] = 0
	return sum
}
