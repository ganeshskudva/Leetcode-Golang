package Medium

func findMaxLength(nums []int) int {
	mp, sum, max := make(map[int]int), 0, 0
	mp[0] = -1

	for i := range nums {
		if nums[i] == 0 {
			sum++
		} else {
			sum--
		}

		if val, ok := mp[sum]; ok {
			if (i - val) > max {
				max = i - val
			}
		} else {
			mp[sum] = i
		}
	}

	return max
}
