package Medium

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	if len(nums) == 0 {
		return ret
	}
	ret[0] = bs(nums, target, 0, true)
	if ret[0] == -1 {
		return ret
	}
	ret[1] = bs(nums, target, ret[0], false)
	
	return ret
}

func bs(nums []int, tgt, start int, first bool) int {
	l, h, mid, ret := start, len(nums)-1, 0, -1
	
	for l <= h {
		mid = l + (h-l)/2
		if nums[mid] < tgt {
			l = mid + 1
		} else if nums[mid] > tgt {
			h = mid -1
		} else {
			if nums[mid] == tgt {
				ret = mid
			}
			if first {
				h = mid -1
			} else {
				l = mid+1
			}
		}
	}
	
	return ret
}
