package Hard

func numberOfArithmeticSlices(nums []int) int {
	n, res := len(nums), 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			mp := make(map[string]int)
			res += solve(nums, j + 1, int64(nums[j]) - int64(nums[i]), 2, mp)
		}
	}
	
	return res
}

func solve(nums []int, idx int, diff int64, cnt int, mp map[string]int) int {
	res, key := 0, fmt.Sprintf("%d-%d", idx, diff)
	if v, ok := mp[key]; ok {
		return v
	}
	
	if cnt >= 3 {
		res++ 
	}
	
	for i := idx; i < len(nums); i++ {
		if int64(nums[i]) - int64(nums[idx - 1]) == diff {
			res += solve(nums, i + 1, diff, cnt + 1, mp)
		}
	}
	mp[key] = res
	return res
}
