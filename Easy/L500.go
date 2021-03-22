package Easy

func findWords(words []string) []string {
	var ans []string
	if len(words) == 0 {
		return ans
	}

	first := make(map[byte]bool)
	sec := make(map[byte]bool)
	third := make(map[byte]bool)

	f := []byte{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}
	initMap(f, first)

	f = []byte{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}
	initMap(f, sec)

	f = []byte{'z', 'x', 'c', 'v', 'b', 'n', 'm'}
	initMap(f, third)

	lst := []map[byte]bool{first, sec, third}

	for _, w := range words {
		if checkContains(strings.ToLower(w), lst) {
			ans = append(ans, w)
		}
	}

	return ans
}

func checkContains(w string, lst []map[byte]bool) bool {
	i := 0
	for _, l := range lst {
		for i = 0; i < len(w); i++ {
			if _, ok := l[w[i]]; !ok {
				break
			}
		}
		if i == len(w) {
			return true
		}
	}

	return false
}

func initMap(f []byte, mp map[byte]bool) {
	for _, c := range f {
		mp[c] = true
	}
}
