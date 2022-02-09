package Medium

func findPairs(nums []int, k int) int {
	if len(nums) == 0 || k < 0 {
		return 0
	}

	mp, cnt := make(map[int]int), 0
	for _, n := range nums {
		mp[n]++
	}

	for key, val := range mp {
		if k == 0 {
			if val >= 2 {
				cnt++
			}
		} else {
			if _, ok := mp[key+k]; ok {
				cnt++
			}
		}
	}

	return cnt
}
