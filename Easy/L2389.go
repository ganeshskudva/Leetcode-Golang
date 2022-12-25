package Easy

func answerQueries(nums []int, queries []int) []int {
	n, m, sum := len(nums), len(queries), 0
	
	sort.Ints(nums)
	var preSum []int
	
	for i := 0; i < n; i++ {
		sum += nums[i]
		preSum = append(preSum, sum)
	}
	
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if preSum[j] <= queries[i] {
				ans[i] = j + 1
			} else {
				break
			}
		}
	}
	
	return ans
}
