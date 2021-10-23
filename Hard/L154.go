package Hard

func findMin(nums []int) int {
	lo, hi, mid := 0, len(nums)-1, 0

	for lo < hi {
		mid = lo + (hi-lo)/2

		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else if nums[mid] < nums[hi] {
			hi = mid
		} else {
			hi--
		}
	}

	return nums[lo]
}
