package Easy

func findErrorNums(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2)
	mp := make(map[int]bool)

	for _, v := range nums {
		if _, ok := mp[v]; ok {
			ans[0] = v
		}
		mp[v] = true
	}

	for i:= 1; i<= n; i++ {
		if _, ok := mp[i]; !ok {
			ans[1] = i
		}
	}

	return ans
}
