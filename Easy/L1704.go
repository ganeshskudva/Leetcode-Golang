package Easy

func halvesAreAlike(s string) bool {
	n := len(s)
	cntA, cntB := 0, 0
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	mp := make(map[rune]bool)

	for _, v := range vowels {
		mp[v] = true
	}

	for i := 0; i < n/2; i++ {
		if mp[rune(s[i])] {
			cntA++
		}
	}

	for i := n / 2; i < n; i++ {
		if mp[rune(s[i])] {
			cntB++
		}
	}

	return cntA == cntB
}
