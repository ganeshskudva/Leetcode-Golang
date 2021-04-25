package Medium

func rotate(matrix [][]int)  {
  s, e := 0, len(matrix)-1
	for s < e {
		tmp := matrix[s]
		matrix[s] = matrix[e]
		matrix[e] = tmp
		s++
		e--
	}
	
	for i:= 0; i < len(matrix); i++ {
		for j := i+1; j < len(matrix[i]); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}
