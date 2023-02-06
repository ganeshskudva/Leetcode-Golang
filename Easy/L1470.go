package Easy

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i, j, idx := 0, n, 0; idx < len(res); i, j = i+1, j+1 {
		res[idx] = nums[i]
		idx++
		res[idx] = nums[j]
		idx++
	}

	return res
}
