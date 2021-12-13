package Hard

func nthMagicalNumber(n int, a int, b int) int {
	A, B, mod := a, b, 1000000007
	l, r := min(a, b), n*min(a, b)

	for B > 0 {
		B, A = A%B, B
	}

	lcm := (a * b) / A
	for l < r {
		mid := l + (r-l)/2

		if (mid/a)+(mid/b)-(mid/lcm) < n {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l % mod
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
