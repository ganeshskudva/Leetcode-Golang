package Easy

func arrangeCoins(n int) int {
	i := 0
	for n > 0 {
		i += 1
		n -= i
	}

	if n == 0 {
		return i
	}

	return i - 1
}
