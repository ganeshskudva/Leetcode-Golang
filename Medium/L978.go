package Medium

func maxTurbulenceSize(arr []int) int {
	inc, dec, result := 1, 1, 1

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			dec = inc + 1
			ince = 1
		} else if arr[i] > arr[i-1] {
			inc = dec + 1
			dec = 1
		} else {
			inc = 1
			dec = 1
		}

		result = max(result, max(dec, inc))
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
