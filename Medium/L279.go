package Medium

func numSquares(n int) int {
	mp := make(map[int]int)

	return count(mp, n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	
	return b
}

func count(mp map[int]int, n int) int {
	maxInt := 1<<(bits.UintSize-1) - 1
	
	if n == 0 {
		return 0
	}
	
	if ret, ok := mp[n]; ok {
		return ret
	}
	
	out := maxInt
	for i := 1; i*i <= n; i++ {
		ret := count(mp, n - i*i)
		out = min(out, ret)
	}
	
	mp[n]= out + 1
	
	return mp[n]
}
