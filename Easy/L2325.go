package Easy

func decodeMessage(key string, message string) string {
	var sb strings.Builder
	hash, j := make(map[byte]byte), 0

	for i := range key {
		if key[i] == ' ' {
			continue
		}
		if _, ok := hash[key[i]]; !ok {
			hash[key[i]] = byte('a' + j)
			j++
		}
	}

	for i := range message {
		if message[i] == ' ' {
			sb.WriteByte(' ')
			continue
		}
		by := hash[message[i]]
		sb.WriteByte(by)
	}

	return sb.String()
}
