package Hard

func atMostNGivenDigitSet(digits []string, n int) int {
	if len(digits) == 0 || n == 0 {
		return 0
	}

	nStr := fmt.Sprintf("%d", n)
	numDigitsN, total, numDigits := len(nStr), 0, len(digits)

	for i := 1; i < numDigitsN; i++ {
		total += IntPow(numDigits, i)
	}

	for i := 0; i < numDigitsN; i++ {
		hasSameNum := false
		for _, d := range digits {
			if d[0] < nStr[i] {
				total += IntPow(numDigits, numDigitsN-i-1)
			} else if d[0] == nStr[i] {
				hasSameNum = true
			}
		}
		if !hasSameNum {
			return total
		}
	}

	return total + 1
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
