package Hard

func trap(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	left, right, sum := leftbound(height), rightbound(height), 0

	for i := range height {
		min := left[i]
		if right[i] < left[i] {
			min = right[i]
		}

		sum += min - height[i]
	}

	return sum
}

func leftbound(arr []int) []int {
	ret := make([]int, len(arr))

	ret[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		ret[i] = ret[i-1]
		if arr[i] > ret[i-1] {
			ret[i] = arr[i]
		}
	}

	return ret
}

func rightbound(arr []int) []int {
	ret := make([]int, len(arr))

	ret[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		ret[i] = ret[i+1]
		if arr[i] > ret[i+1] {
			ret[i] = arr[i]
		}
	}

	return ret
}
