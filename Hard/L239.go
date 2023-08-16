package Hard

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{0}
	}

	var (
		res, q []int
	)
	n := len(nums)

	for i := 0; i < n; i++ {
		for len(q) > 0 && q[0] < (i-k+1) {
			q = q[1:]
		}
		for len(q) > 0 && nums[q[len(q)-1]] < nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}

	return res
}
