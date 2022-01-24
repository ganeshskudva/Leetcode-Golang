package Easy

func detectCapitalUse(word string) bool {
	cnt := 0
	for i := 0; i < len(word); i++ {
		if isUpper(word[i]) {
			cnt++
		}
	}

	if cnt == 0 || cnt == len(word) {
		return true
	}

	if cnt == 1 && isUpper(word[0]) {
		return true
	}
	return false
}

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
