package Medium

func multiply(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	products := make([]byte, n1+n2)
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			d1, d2 := num1[i]-'0', num2[j]-'0'
			products[i+j+1] += d1 * d2
		}
	}

	carry := 0
	for i := len(products) - 1; i >= 0; i-- {
		tmp := (int(products[i]) + carry) % 10
		carry = (int(products[i]) + carry) / 10
		products[i] = byte(tmp)
	}

	for len(products) != 0 && products[0] == '0' {
		products = products[1:]
	}

	if len(products) == 0 {
		return "0"
	}

	return string(products)
}
