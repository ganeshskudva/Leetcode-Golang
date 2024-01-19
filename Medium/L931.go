package Medium

func minFallingPathSum(matrix [][]int) int {
	var (
		solve func(row, col int) int
	)
	cols, mp, res := len(matrix[0]), make(map[string]int), math.MaxInt32

	solve = func(row, col int) int {
		if col >= cols || col < 0 {
			return math.MaxInt32
		}

		key := fmt.Sprintf("%d#%d", row, col)
		if v, ok := mp[key]; ok {
			return v
		}

		if row == len(matrix)-1 && col < cols {
			return matrix[row][col]
		}

		mp[key] = matrix[row][col] + min(solve(row+1, col-1), min(solve(row+1, col), solve(row+1, col+1)))
		return mp[key]
	}

	for i := 0; i < cols; i++ {
		res = min(res, solve(0, i))
	}

	return res
}
