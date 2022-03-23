package Easy

func commonChars(words []string) []string {
	var res []string
	count := make([]int, 26)
	for i := 0; i < 26; i++ {
		count[i] = math.MaxInt32
	}
	for _, w := range words {
		cnt := make([]int, 26)
		for i := range w {
			cnt[w[i]-'a']++
		}
		for i := 0; i < 26; i++ {
			if cnt[i] < count[i] {
				count[i] = cnt[i]
			}
		}
	}

	for ch := 'a'; ch <= 'z'; ch++ {
		for i := count[ch-'a']; i > 0; i-- {
			res = append(res, string(ch))
		}
	}

	return res
}
