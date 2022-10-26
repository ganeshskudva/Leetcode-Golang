package Medium

func checkSubarraySum(nums []int, k int) bool {
	mp := make(map[int]int)
	mp[0] = -1
	runningSum := 0
	for i := 0; i < len(nums); i++ {
		runningSum += nums[i]
		if k != 0 {
			runningSum %= k
		}
		if v, ok := mp[runningSum]; ok {
			if i-v > 1 {
				return true
			}
		} else {
			mp[runningSum] = i
		}
	}
	return false
}
