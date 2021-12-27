package Easy

func findComplement(num int) int {
	i, j := 0, 0

	for i < num {
		i += pow(j)
		j++
	}

	return i - num
}

func pow(n int) int {
	if n == 0 {
		return 1
	}
	result := 2
	for i := 2; i <= n; i++ {
		result *= 2
	}

	return result
}
