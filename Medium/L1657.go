package Medium

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	
	solve := func(s string) []int{
		arr := make([]int, 26)
		
		for _, c := range s {
			arr[c - 'a']++
		}
		
		return arr
	}

	charCnt1, charCnt2 := solve(word1), solve(word2)
	n := len(word1)
	freqCnt1, freqCnt2 := make([]int, n + 1), make([]int, n + 1)
	
	for i := 0; i < 26; i++ {
		if charCnt1[i] > 0 {
			freqCnt1[charCnt1[i]]++
			if charCnt2[i] == 0 {
				return false
			}
		}
		
		if charCnt2[i] > 0 {
			freqCnt2[charCnt2[i]]++
			if charCnt1[i] == 0 {
				return false
			}
		}
	}

	for i := 0; i < n + 1; i++ {
		if freqCnt1[i] != freqCnt2[i] {
			return false
		}
	}

	return true
}
