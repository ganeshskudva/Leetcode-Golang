package Easy

func findSpecialInteger(arr []int) int {
	n, t := len(arr), len(arr)/4

	for i := 0; i < n-t; i++ {
		if arr[i] == arr[i+t] {
			return arr[i]
		}
	}

	return -1
}
