package Easy

func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mp := make(map[int][]int)
	for i := range nums {
		if _, ok := mp[nums[i]]; !ok {
			mp[nums[i]] = append(mp[nums[i]], []int{1, i, i}...)
		} else {
			mp[nums[i]][0]++
			mp[nums[i]][2] = i
		}
	}

	deg, res := math.MinInt32, math.MaxInt32
	for _, v := range mp {
		if v[0] > deg {
			deg, res = v[0], v[2]-v[1]+1
		} else if v[0] == deg {
			res = min(res, v[2]-v[1]+1)
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
