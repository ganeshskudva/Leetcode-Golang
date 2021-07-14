package Medium

func customSortString(order string, str string) string {
	arr := make([]int, 26)
	for i := 0; i < len(str); i++ {
		arr[str[i]-'a']++
	}

	var sb strings.Builder
	for i := 0; i < len(order); i++ {
		for arr[order[i]-'a'] > 0 {
			sb.WriteByte(order[i])
			arr[order[i]-'a']--
		}
	}

	for i := 0; i < len(arr); i++ {
		for arr[i] > 0 {
			sb.WriteByte(byte(i) + 'a')
			arr[i]--
		}
	}

	return sb.String()
}
