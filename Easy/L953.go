package Easy

func isAlienSorted(words []string, order string) bool {
	if len(words) == 0 {
		return true
	}
	
	mp := make(map[byte]int)
	for i, _ := range order {
		mp[order[i]] = i
	}
	
	for i := 1; i < len(words); i++ {
		if !inOrder(words[i-1], words[i], mp) {
			return false
		}
	}
    
    return true
	
}

func inOrder(s1, s2 string, mp map[byte]int) bool {
	m, n := len(s1), len(s2)
	for i:= 0; i < min(m, n); i++ {
		if s1[i] != s2[i] {
			if mp[s1[i]] < mp[s2[i]] {
				return true
			} else {
				return false
			}
		}
	}
    
    return m<=n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	
	return b
}
