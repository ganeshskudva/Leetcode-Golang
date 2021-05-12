package Medium

type NumMatrix struct {
	mt [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	mat := make([][]int, m)
	for i := range mat {
		mat[i] = make([]int, n)
		for j := range mat[i] {
			mat[i][j] = 0
		}
	}

	for i := 0; i < m; i++ {
		if n > 1 {
			for j := 1; j < n; j++ {
				if j == 1 {
					mat[i][0] = matrix[i][0]
				}
				mat[i][j] = mat[i][j-1] + matrix[i][j]

			}
		} else {
			mat[i][0] = matrix[i][0]
		}
	}

	return NumMatrix{
		mt: mat,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.mt[i][col2]
		if col1 > 0 {
			sum -= this.mt[i][col1-1]
		}
	}

	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
