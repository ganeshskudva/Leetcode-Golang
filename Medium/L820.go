package Medium

func minimumLengthEncoding(words []string) int {
    	if len(words) == 0 {
		return 0
	}

	mp := make(map[string]bool)
	for _, w := range words {
		mp[w] = true
	}

	for k, _ := range mp {
		for i := 1; i < len(k); i++ {
			s := k[i:]
			delete(mp, s)
		}
	}
	sum := 0
	for w, _ := range mp {
		sum += len(w) + 1
	}
	
	return sum
}
