func findDuplicates(nums []int) []int {
	var ret []int
	if len(nums) == 0 {
		return ret
	}

	for _, n := range nums {
		index := abs(n) - 1
		if nums[index] < 0 {
			ret = append(ret, abs(n))
		}
		nums[index] = -nums[index]
	}

	return ret
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
