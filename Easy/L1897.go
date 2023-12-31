package Easy

func makeEqual(words []string) bool {
	counts := make(map[byte]int)
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			counts[w[i] - 'a']++
		}
	}
	
	n := len(words)
	for _, v := range counts {
		if v % n != 0 {
			return false
		}
	}
	
	return true
}
