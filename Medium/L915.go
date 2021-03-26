package Medium 

func wordSubsets(A []string, B []string) []string {
	var res []string
	if len(A) == 0 || len(B) == 0 {
		return res
	}
	count := make(map[rune]int)

	for _, s := range B {
		tmp := getMap(s)

		for k, v := range tmp {
			count[k] = max(count[k], v)
		}
	}

	for _, s := range A {
		tmp := getMap(s)
		exists := true
		for k, v := range count {
			if tmp[k] < v {
				exists = false
				break
			}
		}

		if exists {
			res = append(res, s)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getMap(s string) map[rune]int {
	mp := make(map[rune]int)

	for _, c := range s {
		mp[c]++
	}

	return mp
}
