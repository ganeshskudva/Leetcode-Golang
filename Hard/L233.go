package Hard

func findIntegers(n int) int {
	f := make([]int, 32)
	f[0], f[1] = 1, 2

	for i := 2; i < 32; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	ans, k, preBit := 0, 30, 0
	for k >= 0 {
		if (n & (1 << k)) > 0 {
			ans += f[k]
			if preBit > 0 {
				return ans
			}
			preBit = 1
		} else {
			preBit = 0
		}
		k--
	}

	return ans + 1
}
