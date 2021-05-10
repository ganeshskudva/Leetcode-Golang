package Easy

func countPrimes(n int) int {
    	mp, cnt := make([]bool, n), 0

	for i := 2; i < n; i++ {
		if !mp[i] {
			cnt++
			for j := 2; i*j < n; j++ {
				mp[i*j] = true
			}
		}
	}

	return cnt
}
