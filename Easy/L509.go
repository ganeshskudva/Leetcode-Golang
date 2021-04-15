package Easy

func fib(n int) int {
	mp := make(map[int]int)
	
	mp[0], mp[1], mp[2] = 0, 1, 1
	
	return fibo(mp, n)
}

func fibo(mp map[int]int, n int) int {
	if _, ok := mp[n]; ok {
		return mp[n]
	}
	
	mp[n] = fibo(mp, n-1) + fibo(mp, n-2)
	
	return mp[n]
}
