package Medium

func minOperations(nums []int, x int) int {
	tgt := -x
	for _, n := range nums {
		tgt += n
	}

	if tgt == 0 {
		return len(nums)
	}

	mp, sum, res := make(map[int]int), 0, math.MinInt32
	mp[0] = -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := mp[sum-tgt]; ok {
			res = maximum(res, i-mp[sum-tgt])
		}

		mp[sum] = i
	}

	if res == math.MinInt32 {
		return -1
	}

	return len(nums) - res
}

func maximum(i, j int) int {
	if i > j {
		return i
	}

	return j
}
