package Easy

func average(salary []int) float64 {
	var sum float64
	minimum, maximum := math.MaxInt32, math.MinInt32

	for i := 0; i < len(salary); i++ {
		sum += float64(salary[i])
		minimum, maximum = min(minimum, salary[i]), max(maximum, salary[i])
	}

	sum -= float64(minimum + maximum)
	return sum / float64(len(salary)-2)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
