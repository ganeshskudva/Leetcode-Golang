package Hard

func winnerSquareGame(n int) bool {
	return solve(n, make(map[int]bool))
}

func solve(n int, mp map[int]bool) bool {
	if n == 0 {
		return false
	}

	if v, ok := mp[n]; ok {
		return v
	}

	res := false
	for i := 1; i*i <= n; i++ {
		if !solve(n-i*i, mp) {
			res = true
			break
		}
	}
	mp[n] = res
	return res
}
