package Medium

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func minSteps(s string, t string) int {
	cnt, arr := 0, make([]int, 26)

	for _, c := range s {
		arr[c-'a']++
	}
	for _, c := range t {
		arr[c-'a']--
	}

	for _, v := range arr {
		cnt += abs(v)
	}

	return cnt
}
