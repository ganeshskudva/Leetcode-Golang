package Medium

func findMatrix(nums []int) [][]int {
	var ans [][]int
	freq := make([]int, len(nums) + 1)
	
	for _, n := range nums {
		if freq[n] >= len(ans) {
			ans = append(ans, []int{n})
		} else {
			ans[freq[n]] = append(ans[freq[n]], n)
		}
		freq[n]++
	}
	
	return ans
}
