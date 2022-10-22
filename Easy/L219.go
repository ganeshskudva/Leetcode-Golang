package Easy

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)

	for i := range nums {
		if v, ok := mp[nums[i]]; ok {
			if i-v <= k {
				return true
			}
		}
		mp[nums[i]] = i
	}

	return false
}
