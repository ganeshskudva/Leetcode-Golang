package Easy

func minStartValue(nums []int) int {
	sum, minSum := 0, 0

	for _, n := range nums {
		sum += n
		if sum < minSum {
			minSum = sum
		}
	}

	return 1 - minSum
}
