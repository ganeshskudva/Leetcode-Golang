package Easy

func hammingDistance(x int, y int) int {
	xor, res := x^y, 0

	for xor != 0 {
		res += xor & 1
		xor >>= 1
	}

	return res
}
