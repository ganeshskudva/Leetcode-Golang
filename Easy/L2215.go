package Easy

var exists struct{}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var res [][]int
	st1, st2, i := make(map[int]struct{}), make(map[int]struct{}), 0

	for i = 0; i < len(nums1); i++ {
		st1[nums1[i]] = exists
		if i < len(nums2) {
			st2[nums2[i]] = exists
		}
	}
	for ; i < len(nums2); i++ {
		st2[nums2[i]] = exists
	}

	var res1, res2 []int
	for k, _ := range st1 {
		if _, ok := st2[k]; ok {
			continue
		}
		res1 = append(res1, k)
	}
	for k, _ := range st2 {
		if _, ok := st1[k]; ok {
			continue
		}
		res2 = append(res2, k)
	}

	res = append(res, res1, res2)
	return res
}
