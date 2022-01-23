package Medium

func sequentialDigits(low int, high int) []int {
	var q []int
	for i := 1; i <= 9; i++ {
		q = append(q, i)
	}

	var ret []int
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n <= high && n >= low {
			ret = append(ret, n)
		}

		if n > high {
			break
		}
		num := n % 10
		if num < 9 {
			q = append(q, n*10+(num+1))
		}
	}

	return ret
}
