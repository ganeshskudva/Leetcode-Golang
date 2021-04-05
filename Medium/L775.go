package Medium

func isIdealPermutation(A []int) bool {
	for i := 0; i < len(A); i++ {
		if math.Abs(float64(A[i])-float64(i)) > 1 {
			return false
		}
	}

	return true
}
