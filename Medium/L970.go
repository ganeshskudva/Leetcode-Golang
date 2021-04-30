package Medium

func powerfulIntegers(x int, y int, bound int) []int {
	mp := make(map[int]bool)
	for i := 1; i < bound; i *= x {
		for j := 1; i+j <= bound; j *= y {
			mp[i+j] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}

	var res []int
	for v, _ := range mp {
		if mp[v] {
			res = append(res, v)
		}
	}

	return res
}
