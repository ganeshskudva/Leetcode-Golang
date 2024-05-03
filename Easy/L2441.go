package Easy

func findMaxK(nums []int) int {
	mp, res := make(map[int]bool), -1
	
	for _, n := range nums {
		if mp[-n] {
			res = max(res, abs(n))
		}
		mp[n] = true
	}
	
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	
	return a
}
