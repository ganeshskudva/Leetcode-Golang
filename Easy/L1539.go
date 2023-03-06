package Easy

func findKthPositive(arr []int, k int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid]-mid <= k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left + k
}
