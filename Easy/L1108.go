package easy

func defangIPaddr(address string) string {
	var s strings.Builder

	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			s.WriteString("[.]")
		} else {
			s.WriteByte(address[i])
		}
	}

	return s.String()
}
