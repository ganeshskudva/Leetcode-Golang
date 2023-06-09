package Easy

func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters) - 1
	
	if target >= letters[len(letters) - 1] {
		return letters[0]
	}
	
	for low < high {
		mid := low + (high - low) / 2
		if letters[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	
	return letters[low % len(letters)]
}
