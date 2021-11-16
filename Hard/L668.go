package Hard

func findKthNumber(m int, n int, k int) int {
	lo, hi, mid := 0, m*n, 0

	for lo < hi {
		mid = lo + (hi-lo)/2

		if kSmaller(n, m, mid, k) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func kSmaller(n, m, num, k int) bool {
	c := 0
	for i := 0; i <= m; i++ {
		c += min(n, num/i)
	}

	return c >= k
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
