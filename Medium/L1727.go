package Medium

func largestSubmatrix(matrix [][]int) int {
	rn, cn := len(matrix), len(matrix[0])

	for i := 1; i < rn; i++ {
		for j := 0; j < cn; j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}
	
	res := 0
	for i := 0; i < rn; i++ {
		sort.Ints(matrix[i])
		for j := 0; j < cn; j++ {
			res = max(res, matrix[i][j] * (cn - j))
		}
	}
	
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
