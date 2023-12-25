package Easy

func isPathCrossing(path string) bool {
	var (
		getNextPoint func(b byte) ([]int, string)
		key          string
	)
	visited, curr := make(map[string]bool), []int{0, 0}
	visited[fmt.Sprintf("%d#%d", curr[0], curr[1])] = true

	getNextPoint = func(b byte) ([]int, string) {
		mp := map[byte][]int{
			'N': {0, 1},
			'S': {0, -1},
			'E': {1, 1},
			'W': {-1, -1},
		}

		nextPoint := []int{curr[0] + mp[b][0], curr[1] + mp[b][1]}
		return nextPoint, fmt.Sprintf("%d#%d", nextPoint[0], nextPoint[1])
	}

	for _, p := range path {
		curr, key = getNextPoint(byte(p))

		if _, ok := visited[key]; ok {
			return true
		}
		visited[key] = true
	}

	return false
}
