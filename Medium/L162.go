package Medium

func findPeakElement(nums []int) int {
	l, h, mid := 0, len(nums)-1, 0

	for l < h {
		mid = l + (h-l)/2

		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			h = mid
		}
	}

	return l
}
