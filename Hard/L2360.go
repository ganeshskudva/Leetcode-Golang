package Hard

func longestCycle(edges []int) int {
	var exists struct{}
	res, st := -1, make(map[int]struct{})

	for _, e := range edges {
		if e != -1 {
			if _, contains := st[e]; !contains {
				mp, pos := make(map[int]int), 0
				mp[e], st[e] = pos, exists

				for edges[e] != -1 {
					e = edges[e]

					if _, contains = mp[e]; contains {
						res = max(res, pos-mp[e]+1)
						break
					}

					if _, contains = st[e]; contains {
						break
					}

					pos++
					mp[e] = pos
					st[e] = exists
				}
			}
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
