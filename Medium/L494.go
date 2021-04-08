package Medium

type Entry struct {
	index, sum int
}

func findTargetSumWays(nums []int, S int) int {
	mp := make(map[Entry]int)
	
	return dp(nums, S, len(nums)-1, 0, mp)
}

func dp(nums []int, target, index, curSum int, mp map[Entry]int) int {
	e := Entry{
		index: index,
		sum:   curSum,
	}
	if _, ok := mp[e]; ok {
		return mp[e]
	}
	
	if index < 0 && curSum == target {
		return 1
	}
	if index < 0 {
		return 0
	}
	
	pos := dp(nums, target, index - 1, nums[index]+curSum, mp)
	neg := dp(nums, target, index - 1, -nums[index]+curSum, mp)
	
	e = Entry{
		index: index,
		sum:   curSum,
	}
	mp[e] = pos+neg
	
	return pos+neg
}
