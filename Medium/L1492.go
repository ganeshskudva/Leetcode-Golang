package Medium

func kthFactor(n int, k int) int {
	root := math.Sqrt(float64(n))

	for i := 1; i < int(math.Ceil(root)); i++ {
		if n%i == 0 {
			k--
			if k == 0 {
				return i
			}
		}
	}

	for i := int(root); i > 0; i-- {
		if n%i == 0 {
			k--
			if k == 0 {
				return n / i
			}
		}
	}

	return -1
}
