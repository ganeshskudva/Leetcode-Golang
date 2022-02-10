package Medium

func subarraySum(nums []int, k int) int {
    	sum, res, mp := 0, 0, make(map[int]int)
	mp[0] = 1
	
	for i := range nums {
		sum += nums[i]
		if val, ok := mp[sum-k]; ok {
			res += val
		}
		mp[sum]++
	}
	
	return res
}
