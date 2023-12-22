package Easy

func maxScore(s string) int {
	zeroes, ones, maximum := 0, 0, math.MinInt32

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeroes++
		} else {
			ones++
		}

		if i != len(s)-1 {
			maximum = max(maximum, zeroes-ones)
		}
	}

	return maximum + ones
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
