package Medium

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	low, high := 0, len(nums)-1
	for i := low; i <= high; {
		if nums[i] == 0 {
			nums[i], nums[low] = nums[low], nums[i]
			i++
			low++
		} else if nums[i] == 2 {
			nums[i], nums[high] = nums[high], nums[i]
			high--
		} else {
			i++
		}
	}
}
