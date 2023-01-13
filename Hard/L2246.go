package Hard

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestPath(parent []int, s string) int {
	longstPathVal, mp := math.MinInt32, make(map[int][]int)

	for i := range parent {
		mp[parent[i]] = append(mp[parent[i]], i)
	}

	solve(0, parent, mp, s, &longstPathVal)
	if longstPathVal == math.MinInt32 {
		return 1
	}

	return longstPathVal
}

func solve(node int, parent []int, mp map[int][]int, s string, longstPathVal *int) int {
	if _, ok := mp[node]; !ok {
		return 1
	}
	maxi, secondMaxi := 0, 0

	for _, nei := range mp[node] {
		longestPathFromNei := solve(nei, parent, mp, s, longstPathVal)
		if s[node] == s[nei] {
			continue
		}
		if longestPathFromNei > maxi {
			secondMaxi = maxi
			maxi = longestPathFromNei
		} else if longestPathFromNei > secondMaxi {
			secondMaxi = longestPathFromNei
		}
	}

	*longstPathVal = max(*longstPathVal, maxi+secondMaxi+1)
	return maxi + 1
}
