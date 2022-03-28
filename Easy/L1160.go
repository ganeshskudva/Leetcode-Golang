package Easy

func countCharacters(words []string, chars string) int {
    	res, cnt := 0, make([]int, 26)

	for i := range chars {
		cnt[chars[i]-'a']++
	}

	for _, w := range words {
		match, tmp := true, make([]int, 26)
		copy(tmp, cnt)

		for i := range w {
			tmp[w[i]-'a']--
			if tmp[w[i]-'a'] < 0 {
				match = false
				break
			}
		}

		if match {
			res += len(w)
		}
	}

	return res
}
