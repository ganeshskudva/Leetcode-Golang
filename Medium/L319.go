package Medium

func bulbSwitch(n int) int {
	low, high := 0, n
	
	for low <= high {
		mid := low + (high - low) / 2
		candidate := mid * mid
		if candidate == n {
			return mid
		} else if candidate < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	
	return low - 1
}
