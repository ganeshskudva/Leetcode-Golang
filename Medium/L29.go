package Medium

func divide(dividend int, divisor int) int {
    if dividend == -2147483648 && divisor == 1 {
		return -2147483648
	}
	n := dividend
	d := divisor
	if dividend < 0 {
		n = -dividend
	}
	if divisor < 0 {
		d = -divisor
	}

	bits := 0
	for i := n; i > 0; i = i >> 1 {
		bits++
	}

	q := 0
	r := 0
	for i := bits - 1;  i>= 0; i-- {
		r = r << 1
		r += (n >> i) & 0x1
		if r >= d {
			r = r - d
			q += 1 << i
		}
	}

	if q > 0x7FFFFFFF {
		q = 0x7FFFFFFF
	}
	if (dividend < 0 && divisor > 0) || (dividend >= 0 && divisor < 0) {
		return -q
	}
	return q
}
