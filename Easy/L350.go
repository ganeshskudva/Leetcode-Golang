package Easy

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	mp1, mp2 := make(map[int]int), make(map[int]int)

	for _, n := range nums1 {
		mp1[n]++
	}

	for _, n := range nums2 {
		mp2[n]++
	}

	for k, v := range mp2 {
		if val, ok := mp1[k]; ok {
			minVal := v
			if val < minVal {
				minVal = val
			}
			for i := 0; i < minVal; i++ {
				res = append(res, k)
			}
		}
	}

	return res
}
