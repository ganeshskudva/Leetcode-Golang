package Medium

func leastBricks(wall [][]int) int {
	if len(wall) == 0 {
		return 0
	}
	
	cnt := 0
	mp := make(map[int]int)
	
	for _, w := range wall {
		sum := 0
		for i:= 0; i < len(w)-1; i++ {
			sum += w[i]
			mp[sum]++
			cnt = max(mp[sum], cnt)
		}
	}
	
	return len(wall) - cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
