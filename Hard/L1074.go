package Hard

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	res, m, n := 0, len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
            mp := make(map[int]int)
			mp[0] = 1
			cur := 0
			for k := 0; k < m; k++ {
				if i > 0 {
					cur += matrix[k][j] - matrix[k][i - 1]
				} else {
					cur += matrix[k][j]
				}
				if v, ok := mp[cur - target]; ok {
					res += v
				}
				mp[cur]++ 
			}
		}
	}
	
	return res
}
