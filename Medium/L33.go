package Medium

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var rotationCnt func() int

	rotationCnt = func() int {
		l, r := 0, len(nums)-1

		for l < r {
			mid := l + (r-l)/2
			if mid < r && nums[mid] > nums[mid+1] {
				return mid + 1
			}
			if mid > l && nums[mid] < nums[mid-1] {
				return mid
			}
			if nums[mid] < nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		return l
	}
	rotCnt := rotationCnt()
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		realMid := (mid + rotCnt) % len(nums)
		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
