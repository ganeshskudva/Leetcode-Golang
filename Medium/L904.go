package Medium

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func totalFruit(fruits []int) int {
	mp := make(map[int]int)
	
	i, j, res := 0, 0, 0
	
	for j < len(fruits) {
		mp[fruits[j]]++
		
		if len(mp) <= 2 {
			res = max(res, j - i + 1)
		} else {
			mp[fruits[i]]--
			if mp[fruits[i]] == 0 {
				delete(mp, fruits[i])
			}
			i++
		}
		j++
	}
	
	return res
}
