package Medium

func openLock(deadends []string, target string) int {
	var (
		neighbors func(code string) []string
		q         []string
	)
	deadSet := make(map[string]bool)
	for _, dead := range deadends {
		deadSet[dead] = true
	}
	if deadSet["0000"] {
		return -1
	}

	neighbors = func(code string) []string {
		var res []string
		for i := 0; i < 4; i++ {
			x := int(code[i] - '0')
			for diff := -1; diff <= 1; diff += 2 {
				y := (x + diff + 10) % 10
				res = append(res, fmt.Sprintf("%s\"\"%d%s", code[:i], y, code[i+1:]))
			}
		}

		return res
	}

	q = append(q, "0000")
	for steps := 0; len(q) > 0; steps++ {
		for i := len(q); i > 0; i-- {
			curr := q[0]
			q = q[1:]
			if curr == target {
				return steps
			}
			for _, nei := range neighbors(curr) {
				if deadSet[nei] {
					continue
				}
				deadSet[nei] = true
				q = append(q, nei)
			}
		}
	}

	return -1
}
