package Medium

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mp, n, res := make(map[int]int), len(nums1), 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := nums3[i] + nums4[j]
			mp[sum]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := nums1[i] + nums2[j]

			if val, ok := mp[-sum]; ok {
				res += val
			}
		}
	}

	return res
}
