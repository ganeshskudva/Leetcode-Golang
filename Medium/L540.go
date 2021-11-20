package Medium

func singleNonDuplicate(nums []int) int {
	lo, hi, mid := 0, len(nums)-1, -1

	for lo < hi {
		mid = lo + (hi-lo)/2

		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				lo = mid + 2
			} else {
				hi = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	}

	return nums[lo]
}
