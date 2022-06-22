package Medium

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, low, high, k int) int {
	pivot := low

	j := low
	for ; j < high; j++ {
		if nums[j] <= nums[high] {
			nums[pivot], nums[j] = nums[j], nums[pivot]
			pivot++
		}
	}

	nums[pivot], nums[j] = nums[j], nums[pivot]
	cnt := high - pivot + 1

	if cnt == k {
		return nums[pivot]
	} else if cnt > k {
		return quickSelect(nums, pivot+1, high, k)
	}

	return quickSelect(nums, low, pivot-1, k-cnt)
}
