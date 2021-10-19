package Easy

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp, res := make(map[int]int), make([]int, len(nums1))
	var st []int
	for _, n := range nums2 {
		for len(st) > 0 && st[len(st)-1] < n {
			mp[st[len(st)-1]] = n
			st = st[:len(st)-1]
		}
		st = append(st, n)
	}

	for i := range nums1 {
		val, ok := mp[nums1[i]]
		if !ok {
			res[i] = -1
		} else {
			res[i] = val
		}
	}

	return res
}
