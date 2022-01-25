package Easy

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	start, end := 0, len(arr)-1
	for start < end {
		if arr[start+1] > arr[start] {
			start++
		} else if arr[end-1] > arr[end] {
			end--
		} else {
			break
		}
	}

	return start != 0 && end != len(arr)-1 && start == end
}
