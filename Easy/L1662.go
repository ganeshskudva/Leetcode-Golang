package Easy

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	idx1, idx2, arrIdx1, arrIdx2 := 0, 0, 0, 0

	for arrIdx1 < len(word1) && arrIdx2 < len(word2) {
		if word1[arrIdx1][idx1] != word2[arrIdx2][idx2] {
			return false
		}

		if idx1 == len(word1[arrIdx1])-1 {
			idx1 = 0
			arrIdx1++
		} else {
			idx1++
		}

		if idx2 == len(word2[arrIdx2])-1 {
			idx2 = 0
			arrIdx2++
		} else {
			idx2++
		}
	}

	return arrIdx1 == len(word1) && arrIdx2 == len(word2)
}
