package Hard

func minDistance(word1 string, word2 string) int {
	m, n, mp := len(word1), len(word2), make(map[string]int)
	return solve(word1, word2, m, n, mp)
}

func solve(w1, w2 string, m, n int, mp map[string]int) int {
	key := fmt.Sprintf("%d-%d", m, n)
	if m == 0 {
		mp[key] = n
		return mp[key]
	}

	if n == 0 {
		mp[key] = m
		return mp[key]
	}

	if _, ok := mp[key]; ok {
		return mp[key]
	}

	if w1[m-1] == w2[n-1] {
		return solve(w1, w2, m-1, n-1, mp)
	}

	insertChar := solve(w1, w2, m-1, n, mp)
	deleteChar := solve(w1, w2, m, n-1, mp)
	replaceChar := solve(w1, w2, m - 1, n - 1, mp)
	
	mp[key] = 1 + min(insertChar, min(replaceChar, deleteChar))
	return mp[key]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	
	return b
}
