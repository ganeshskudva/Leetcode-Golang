package Medium

func numTilings(n int) int {
	return solve(n, 2, make(map[string]int))
}

func solve(n, colState int, memo map[string]int) int {
	mod, ways := 1000000007, 0
	if n < 0 {
		return 0
	}

	if n == 0 {
		if colState == 2 {
			return 1
		} else {
			return 0
		}
	}

	key := fmt.Sprintf("%d|%d", n, colState)
	if val, ok := memo[key]; ok {
		return val
	}

	switch colState {
	case 0:
		ways += solve(n-2, 2, memo)
		ways += solve(n-1, 1, memo)
	case 1:
		ways += solve(n-2, 2, memo)
		ways += solve(n-1, 0, memo)
	case 2:
		ways += solve(n-1, 0, memo)
		ways += solve(n-1, 1, memo)
		ways += solve(n-1, 2, memo)
		ways += solve(n-2, 2, memo)
	}

	ways %= mod
	memo[key] = ways
	return ways
}
