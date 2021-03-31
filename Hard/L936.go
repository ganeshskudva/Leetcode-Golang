package Hard

func movesToStamp(stamp string, target string) []int {
	var res []int
	visited := make([]bool, len(target))
	stars := 0

	for stars < len(target) {
		done := false
		for i := 0; i <= len(target)-len(stamp); i++ {
			if !visited[i] && canReplace(stamp, target, i) {
				by := []byte(target)
				stars = doReplace(&by, i, len(stamp), stars)
				target = string(by)
				done = true
				visited[i] = true
				res = append(res, i)
				if stars == len(target) {
					break
				}
			}
		}

		if !done {
			return []int{}
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func canReplace(stamp, target string, p int) bool {
	for i := 0; i < len(stamp); i++ {
		if target[i+p] != '*' && target[i+p] != stamp[i] {
			return false
		}
	}

	return true
}

func doReplace(target *[]byte, p, length, count int) int {
	for i := 0; i < length; i++ {
		if (*target)[i+p] != '*' {
			(*target)[i+p] = '*'
			count++
		}
	}

	return count
}
