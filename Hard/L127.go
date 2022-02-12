package Hard

var exists struct{}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	beginSet, endSet, visited, dict := make(map[string]struct{}), make(map[string]struct{}), make(map[string]struct{}), make(map[string]struct{})

	for _, w := range wordList {
		dict[w] = exists
	}

	if _, ok := dict[endWord]; !ok {
		return 0
	}

	beginSet[beginWord], endSet[endWord] = exists, exists
	length := 1

	for len(beginSet) > 0 && len(endSet) > 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

		tmp := make(map[string]struct{})
		for s, _ := range beginSet {
			word := []byte(s)
			for i := 0; i < len(word); i++ {
				orig := word[i]
				for ch := 'a'; ch <= 'z'; ch++ {
					if byte(ch) == orig {
						continue
					}
					word[i] = byte(ch)
					target := string(word)
					if _, ok := endSet[target]; ok {
						return length + 1
					}
					if _, ok := visited[target]; !ok {
						if _, contains := dict[target]; contains {
							visited[target] = exists
							tmp[target] = exists
						}
					}
					word[i] = orig
				}
			}
		}
		length, beginSet = length+1, tmp
	}

	return 0
}
