package Medium

func minFlips(a int, b int, c int) int {
	flips := 0
	for a > 0 || b > 0 || c > 0 {
		bitA, bitB, bitC := a&1, b&1, c&1

		if bitC == 0 {
			flips += bitA + bitB
		} else if bitA == 0 && bitB == 0 {
			flips++
		}

		a, b, c = a>>1, b>>1, c>>1
	}

	return flips
}
