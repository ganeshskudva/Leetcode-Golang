package Medium

func minSpeedOnTime(dist []int, hour float64) int {
	var timeTaken func(speed int) float64
	left, right, res := 1, 10000007, -1

	timeTaken = func(speed int) float64 {
		var ans float64

		for i := 0; i < len(dist); i++ {
			if i == len(dist)-1 {
				ans += float64(dist[i]) / float64(speed)
			} else {
				ans += math.Ceil(float64(dist[i]) / float64(speed))
			}
		}

		return ans
	}

	for left <= right {
		mid := right - (right-left)/2
		dur := timeTaken(mid)

		if dur <= hour {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}
