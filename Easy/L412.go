func fizzBuzz(n int) []string {
	var res []string
	if n == 0 {
		return res
	}

	for i := 1; i <= n; i++ {
		if isDiv3(i) && isDiv5(i) {
			res = append(res, "FizzBuzz")
		} else if isDiv3(i) && !isDiv5(i) {
			res = append(res, "Fizz")
		} else if !isDiv3(i) && isDiv5(i) {
			res = append(res, "Buzz")
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}

	return res
}

func isDiv3(n int) bool {
	return n%3 == 0
}

func isDiv5(n int) bool {
	return n%5 == 0
}
