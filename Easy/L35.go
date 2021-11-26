package Easy

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	l, r, mid := 0, len(nums)-1, 0

	for l <= r {
		mid = (r-l)/2 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}
