package Hard

func minPatches(nums []int, n int) int {
	miss, added, i := 1, 0, 0

	for miss <= n {
		if i < len(nums) && nums[i] <= miss {
			miss += nums[i]
			i++
		} else {
			miss += miss
			added++
		}
	}

	return added
}
