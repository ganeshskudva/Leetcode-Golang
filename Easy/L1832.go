package Easy

var exists struct{}

func checkIfPangram(sentence string) bool {
	mp := make(map[byte]struct{})

	for i := 0; i < len(sentence); i++ {
		mp[sentence[i]] = exists
	}

	return len(mp) == 26
}
