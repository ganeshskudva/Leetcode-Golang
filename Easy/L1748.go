package Easy

func sumOfUnique(nums []int) int {
    	if len(nums) == 0 {
		return 0
	}

	mp := make(map[int]bool)

	sum := 0
	for _, v := range nums {
		if _, ok := mp[v]; ok {
			if mp[v] {
				sum -= v
				mp[v] = false
			}
		} else {
			mp[v] = true
			sum += v
		}
	}

	return sum
}
