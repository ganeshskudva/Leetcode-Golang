package Easy

func countOdds(low int, high int) int {
	cnt := (high - low) / 2

	if isOdd(low) || isOdd(high) {
		cnt++
	}

	return cnt
}

func isOdd(n int) bool {
	return n%2 != 0
}
