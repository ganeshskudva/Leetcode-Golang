package Easy

func addStrings(num1 string, num2 string) string {
	carry, i, j := 0, len(num1)-1, len(num2)-1
	var sb strings.Builder

	for i >= 0 || j >= 0 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		sb.WriteByte(byte(sum%10 + '0'))
		i--
		j--
	}

	if carry != 0 {
		sb.WriteByte(byte(carry + '0'))
	}

	return reverse(sb.String())
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
