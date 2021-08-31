package Medium

func findMin(nums []int) int {
	l, h, mid, n := 0, len(nums)-1, 0, len(nums)

	for l < h {
		mid = l + (h-l)/2
		if nums[mid] < nums[h] {
			h = mid
		} else {
			l = (mid + 1) % n
		}
	}

	return nums[l]
}
