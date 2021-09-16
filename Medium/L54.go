package Medium

func spiralOrder(matrix [][]int) []int {
    	var res []int
	if len(matrix) == 0 {
		return res
	}

	n, m := len(matrix), len(matrix[0])
	up, down, left, right := 0, n-1, 0, m-1

	for len(res) < n*m {
		for j := left; j <= right && len(res) < n*m; j++ {
			res = append(res, matrix[up][j])
		}

		for i := up + 1; i <= down-1 && len(res) < n*m; i++ {
			res = append(res, matrix[i][right])
		}
		
		for j := right; j >= left && len(res) < n*m; j-- {
			res = append(res, matrix[down][j])
		}
		
		for i:= down-1; i >= up + 1 && len(res) < n*m; i-- {
			res = append(res, matrix[i][left])
		}
		
		left++
		right--
		up++
		down--
	}
	
	return res
}
