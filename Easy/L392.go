package Easy

func isSubsequence(s string, t string) bool {
    	if len(s) == 0 {
		return true
	}

	indexS, indexT := 0, 0
	for indexT < len(t) {
		if t[indexT] == s[indexS] {
			indexS++
			if indexS == len(s) {
				return true
			}
		}
		indexT++
	}
	
	return false
}
