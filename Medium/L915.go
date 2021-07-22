package Medium

func partitionDisjoint(nums []int) int {
	localMax, partitionIdx := nums[0], 0
	max := localMax

	for i := range nums {
		if localMax > nums[i] {
			localMax = max
			partitionIdx = i
		} else {
			if nums[i] > max {
				max = nums[i]
			}
		}
	}

	return partitionIdx + 1
}
