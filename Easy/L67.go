package Easy

func addBinary(a string, b string) string {
	var sb strings.Builder
	i, j, carry := len(a)-1, len(b)-1, 0

	for i >= 0 || j >= 0 {
		sum := carry
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		sb.WriteString(fmt.Sprintf("%d", sum%2))
		carry = sum / 2
	}

	if carry != 0 {
		sb.WriteString(fmt.Sprintf("%d", carry))
	}

	ret := sb.String()
	return reverse(ret)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
