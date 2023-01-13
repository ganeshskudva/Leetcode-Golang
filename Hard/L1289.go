package Hard

func minFallingPathSum(grid [][]int) int {
	n, mp, res := len(grid), make([][]int, len(grid)), math.MaxInt32
	if n == 1 {
		return grid[0][0]
	}
	
	for i := 0; i < n; i++ {
		mp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mp[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		res = min(res, solve(0, i, grid, &mp))
	}

	return res
}

func solve(r, c int, grid [][]int, mp *[][]int) int {
	if r == len(grid) {
		return 0
	}
	
	if (*mp)[r][c] != -1 {
		return (*mp)[r][c]
	}

	res := math.MaxInt32
	for k := 0; k < len(grid[0]); k++ {
		if c == k {
			continue
		}
		res = min(res, grid[r][c]+solve(r+1, k, grid, mp))
	}

	(*mp)[r][c] = res
	return (*mp)[r][c]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
