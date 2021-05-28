package Medium

func maximumUniqueSubarray(nums []int) int {
    	i, j, sum, res, vis := 0, 0, 0, 0, make([]bool, 10001)
	
	for j < len(nums) {
		for j < len(nums) && !vis[nums[j]] {
			vis[nums[j]] = true
			sum += nums[j]
			j++
		}
		res = max(res, sum)
		
		for j < len(nums) && vis[nums[j]] {
			vis[nums[i]] = false
			sum -= nums[i]
			i++
		}
	}
	
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
