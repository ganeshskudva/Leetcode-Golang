package Medium

var exists struct{}

func minMutation(start string, end string, bank []string) int {
	return dfs(start, end, bank, make(map[string]struct{}))
}

func dfs(start, end string, bank []string, visited map[string]struct{}) int {
	if start == end {
		return 0
	}

	min := math.MaxInt32
	for _, s := range bank {
		diff := 0
		for i := 0; i < len(start); i++ {
			if start[i] != s[i] {
				diff++
			}
			if diff > 1 {
				break
			}
		}

		if _, ok := visited[s]; !ok {
			if diff == 1 {
				visited[s] = exists
				h := dfs(s, end, bank, visited)
				if h >= 0 {
					if 1+h < min {
						min = 1 + h
					}
				}
				delete(visited, s)
			}
		}
	}

	if min == math.MaxInt32 {
		return -1
	}

	return min
}
