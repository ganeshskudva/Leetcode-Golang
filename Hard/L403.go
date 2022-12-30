package Hard

var exists struct{}

func canCross(stones []int) bool {
	st, mp, n := make(map[int]struct{}), make(map[string]bool), len(stones)

	for _, s := range stones {
		st[s] = exists
	}

	return solve(st, mp, stones, n, 0, 1)
}

func solve(st map[int]struct{}, mp map[string]bool, stones []int, n, pos, prevStep int) bool {
	if _, ok := st[pos]; !ok {
		return false
	}
	if pos > stones[n-1] {
		return false
	}
	if pos == stones[n-1] {
		return true
	}

	key := fmt.Sprintf("%d-%d", pos, prevStep)
	if v, ok := mp[key]; ok {
		return v
	}

	if pos == 0 {
		mp[key] = solve(st, mp, stones, n, pos+1, 1)
	} else {
		if prevStep-1 > 0 {
			mp[key] = solve(st, mp, stones, n, pos+(prevStep-1), prevStep-1)
			if mp[key] {
				return mp[key]
			}
		}
		mp[key] = solve(st, mp, stones, n, pos+prevStep, prevStep) ||
			solve(st, mp, stones, n, pos+(prevStep+1), prevStep+1)
	}

	return mp[key]
}
