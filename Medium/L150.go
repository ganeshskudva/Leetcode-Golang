package Medium

func evalRPN(tokens []string) int {
	var st []int

	for _, ch := range tokens {
		if ch == "+" || ch == "*" {
			if len(st) > 0 {
				v1, v2 := getOperands(&st)
				if ch == "+" {
					st = append(st, v1+v2)
				} else {
					st = append(st, v1*v2)
				}
			}
		} else if ch == "-" || ch == "/" {
			v1, v2 := getOperands(&st)
			if ch == "-" {
				st = append(st, v1-v2)
			} else {
				st = append(st, v1/v2)
			}
		} else {
			v, err := strconv.Atoi(ch)
			if err != nil {
				return 0
			}
			st = append(st, v)
		}
	}

	return st[0]
}

func getOperands(st *[]int) (int, int) {
	v1, v2 := (*st)[len(*st)-2], (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-2]

	return v1, v2
}
