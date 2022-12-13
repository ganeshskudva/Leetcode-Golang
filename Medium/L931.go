package Medium

func minFallingPathSum(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	mp, ans := make(map[string]int), math.MaxInt32

	for c := 0; c < cols; c++ {
		ans = min(ans, solve(matrix, rows-1, c, mp))
	}

	return ans
}

func solve(matrix [][]int, r, c int, mp map[string]int) int {
	if r == 0 && c < len(matrix[0]) && c >= 0 {
		return matrix[r][c]
	}
	if c >= len(matrix[0]) || c < 0 {
		return math.MaxInt32
	}
	key := fmt.Sprintf("%d-%d", r, c)
	if v, ok := mp[key]; ok {
		return v
	}

	mp[key] = matrix[r][c] + min(min(solve(matrix, r-1, c+1, mp), solve(matrix, r-1, c, mp)), solve(matrix, r-1, c-1, mp))
	return mp[key]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
