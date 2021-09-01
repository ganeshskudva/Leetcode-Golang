package Medium

func arrayNesting(nums []int) int {
	maxLen, visited := 0, make([]bool, len(nums))
	for _, n := range nums {
		currLen := 0
		for !visited[n] {
			currLen++
			visited[n] = true
			n = nums[n]
		}
		if currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}
