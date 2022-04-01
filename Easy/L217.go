package Easy

var exists struct{}
func containsDuplicate(nums []int) bool {
	mp := make(map[int]struct{}

	for _, n := range nums {
		if _, ok := mp[n]; ok {
			return true
		}
		mp[n] = exists
	}

	return len(mp) != len(nums)
}
