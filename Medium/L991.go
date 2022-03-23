package Medium

func brokenCalc(startValue int, target int) int {
	// res for counting number of operation
	res := 0

	for target > startValue {
		if isOdd(target) {
			// if target is odd we will make it even
			target++
		} else {
			// if target is even divide by 2
			target /= 2
		}
		res++
	}

	return res + startValue - target
}

func isOdd(n int) bool {
	return n%2 > 0
}
