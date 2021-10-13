package Easy

func guessNumber(n int) int {
	i, j := 1, n
	for i < j {
		mid := i + (j-i)/2
		switch guess(mid) {
		case 0:
			return mid
		case 1:
			i = mid + 1
		case -1:
			j = mid
		}
	}

	return i
}
