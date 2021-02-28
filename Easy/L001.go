package Easy

func twoSum(nums []int, target int) []int {
    var ans []int
	if len(nums) == 0 {
		return ans
	}

	mp := make(map[int]int)
	for i, v:= range nums {
		if j, ok := mp[target-v]; ok {
			return []int{j,i}
		}
        mp[v] = i
	}

	return ans
}
