package Easy

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	res, ind := 0, 0

	for n > 0 {
		digit := n % 2
		flippedDigit := 0
		if digit != 1 {
			flippedDigit = 1
		}
		res += flippedDigit << ind
		n /= 2
		ind++
	}

	return res
}
