package Medium

func findRedundantConnection(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}
	parent := make([]int, 5000)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	for _, e := range edges {
		if find(parent, e[0]) == find(parent, e[1]) {
			return e
		}
		union(e[0], e[1], &parent)
	}

	return []int{}
}

func find(parent []int, x int) int {
	if parent[x] == x {
		return x
	}

	return find(parent, parent[x])
}

func union(p, q int, parent *[]int) {
	x, y := find(*parent, p), find(*parent, q)

	if x == y {
		return
	}

	(*parent)[y] = x
}
