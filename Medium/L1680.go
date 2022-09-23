package Medium

func concatenatedBinary(n int) int {
	mod, res := 1000000007, 0

	for i := 1; i <= n; i++ {
		tmp := i
		for tmp > 0 {
			tmp, res = tmp/2, res*2
		}
		res = (res + i) % mod
	}

	return res
}
