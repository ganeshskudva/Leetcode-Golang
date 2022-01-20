package Medium

func minEatingSpeed(piles []int, h int) int {
	if len(piles) == 0 {
		return 0
	}
	min, max, sum := math.MaxInt32, math.MinInt32, 0
	for _, p := range piles {
		if p < min {
			min = p
		}
		if p > max {
			max = p
		}
		sum += p
	}

	if h == 1 {
		return sum
	}

	lo, high, mid := 1, 1000000000, 0
	for lo < high {
		mid = lo + (high-lo)/2
		if feasible(piles, mid, h) {
			high = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func feasible(arr []int, maxSpeed, h int) bool {
	time := 0

	for _, a := range arr {
		time += int(math.Ceil(float64(a) / float64(maxSpeed)))
	}

	return time <= h
}
