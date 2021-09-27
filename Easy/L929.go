package Easy

func numUniqueEmails(emails []string) int {
	mp := make(map[string]bool)

	for _, s := range emails {
		mp[cleanStrings(s)] = true
	}

	return len(mp)
}

func cleanStrings(s string) string {
	arr := strings.Split(s, "@")
	var sb strings.Builder

	for _, c := range arr[0] {
		if c == '.' {
			continue
		} else if c == '+' {
			break
		} else {
			sb.WriteByte(byte(c))
		}
	}
	sb.WriteString(fmt.Sprintf("@%s", arr[1]))

	return sb.String()
}
