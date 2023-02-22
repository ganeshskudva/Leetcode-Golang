package Medium

func shipWithinDays(weights []int, days int) int {
	if len(weights) == 0 {
		return 0
	}

	maxi, sum := math.MinInt32, 0
	for _, w := range weights {
		maxi = max(maxi, w)
		sum += w
	}

	if days == 1 {
		return sum
	}

	isFeasible := func(maxWeight int) bool {
		cnt, maximum, day := 0, maxWeight, 0
		for _, w := range weights {
			if w <= maximum {
				cnt++
			} else {
				cnt = 0
				maximum = maxWeight
				day++
			}
			maximum -= w
		}

		return day+1 <= days
	}

	low, high := maxi, sum
	for low < high {
		mid := low + (high-low)/2

		if isFeasible(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
