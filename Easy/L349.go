package Easy

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			size := len(res)
			i, j = i+1, j+1
			if size != 0 && res[size-1] == nums1[i-1] {
				continue
			}
			res = append(res, nums1[i-1])
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return res
}
