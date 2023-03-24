package Medium

type pair struct {
	neighbor, edgeWeight int
}

func newPair(x, y int) pair {
	return pair{
		neighbor:   x,
		edgeWeight: y,
	}
}

func minReorder(n int, connections [][]int) int {
	mp, ans := make(map[int][]pair), 0
	var solve func(node, parent int)

	for _, c := range connections {
		if _, ok := mp[c[0]]; !ok {
			mp[c[0]] = make([]pair, 0)
		}
		mp[c[0]] = append(mp[c[0]], newPair(c[1], 1))

		if _, ok := mp[c[1]]; !ok {
			mp[c[1]] = make([]pair, 0)
		}
		mp[c[1]] = append(mp[c[1]], newPair(c[0], 0))
	}

	solve = func(node, parent int) {
		for _, nei := range mp[node] {
			if nei.neighbor != parent {
				ans += nei.edgeWeight
				solve(nei.neighbor, node)
			}
		}
	}

	solve(0, -1)
	return ans
}
