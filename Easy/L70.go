package Easy

func climbStairs(n int) int {
	mp := make(map[int]int)
	
	mp[1], mp[2], mp[3] = 1,2,3
	return dp(mp, n)
}

func dp(mp map[int]int, n int) int {
	if _, ok := mp[n]; ok {
		return mp[n]
	}
	
	mp[n] = dp(mp, n-1) + dp(mp, n-2)
	
	return mp[n]
}
