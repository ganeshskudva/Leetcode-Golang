package Medium

func spellchecker(wordlist []string, queries []string) []string {
	var ans []string
	if len(wordlist) == 0 || len(queries) == 0 {
		return ans
	}

	ans = make([]string, len(queries))
	capital := make(map[string]string)
	vowel := make(map[string]string)
	words := make(map[string]bool)

	for _, w := range wordlist {
		words[w] = true
	}

	for _, w := range wordlist {
		lower := strings.ToLower(w)
		stripVowel := lower
		for _, c := range []string{"a", "e", "i", "o", "u"} {
			stripVowel = strings.ReplaceAll(stripVowel, c, "#")
		}
		if _, ok := capital[lower]; !ok {
			capital[lower] = w
		}
		if _, ok := vowel[stripVowel]; !ok {
			vowel[stripVowel] = w
		}
	}

	for i := 0; i < len(queries); i++ {
		if _, ok := words[queries[i]]; ok {
			ans[i] = queries[i]
			continue
		}
		lower := strings.ToLower(queries[i])
		stripVowel := lower
		for _, c := range []string{"a", "e", "i", "o", "u"} {
			stripVowel = strings.ReplaceAll(stripVowel, c, "#")
		}
		if _, ok := capital[lower]; ok {
			ans[i] = capital[lower]
		} else if _, ok := vowel[stripVowel]; ok {
			ans[i] = vowel[stripVowel]
		} else {
			ans[i] = ""
		}
	}

	return ans
}
