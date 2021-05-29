package Hard

func totalNQueens(n int) int {
	vertical, leftFall, rightFall, res := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), 0
	solve(n, 0, &res, &vertical, &leftFall, &rightFall)
	
	return res
}

func solve(n, row int, res *int, vertical, leftFall, rightFall *[]bool) {
	if row == n {
		*res++
		return
	}
	
	for col := 0; col < n; col++ {
		if (*vertical)[col] || (*leftFall)[row - col + n -1] || (*rightFall)[row+col] {
			continue
		}
		(*vertical)[col] = true
		(*leftFall)[row - col + n - 1] = true
		(*rightFall)[row + col] = true
		
		solve(n, row+1, res, vertical, leftFall, rightFall)

		(*vertical)[col] = false
		(*leftFall)[row - col + n - 1] = false
		(*rightFall)[row + col] = false
	}
}
