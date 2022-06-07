package Easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k, i = k-1, i-1
		} else {
			nums1[k] = nums2[j]
			k, j = k-1, j-1
		}
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k, j = k-1, j-1
	}
}
