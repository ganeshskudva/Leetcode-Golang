package Medium

func permute(nums []int) [][]int {
	var (
		ans [][]int
		tmp []int
		solve func(tmp []int)
	)
	if len(nums) == 0 {
		return ans
	}
	
	mp := make(map[int]bool)
	solve = func(tmp []int) {
		if len(tmp) == len(nums) {
			tmpCopy := make([]int, len(tmp))
			copy(tmpCopy, tmp)
			ans = append(ans, tmpCopy)
			return
		}
		
		for i := 0; i < len(nums); i++ {
			if mp[nums[i]] {
				continue
			}
			tmp = append(tmp, nums[i])
			mp[nums[i]] = true
			solve(tmp)
			delElem := tmp[len(tmp) - 1]
			tmp = tmp[:len(tmp) - 1]
			delete(mp, delElem)
		}
	}
	
	solve(tmp)
	return ans
}
