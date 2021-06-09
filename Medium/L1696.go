package Medium

func maxResult(nums []int, k int) int {
    	var dq []int
	dq = append(dq, 0)
	for i := 1; i < len(nums); i++ {
		if dq[0]+k < i {
			dq = dq[1:]
		}
		nums[i] += nums[dq[0]]
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= nums[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}

	return nums[len(nums)-1]
}
