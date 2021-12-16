package Medium

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adj := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool)
	}
	for _, e := range edges {
		adj[e[0]][e[1]] = true
		adj[e[1]][e[0]] = true
	}

	var leaves []int
	for i := 0; i < n; i++ {
		if len(adj[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	for n > 2 {
		n -= len(leaves)
		var newLeaves []int
		for _, i := range leaves {
			for v, _ := range adj[i] {
				delete(adj[v], i)
				if len(adj[v]) == 1 {
					newLeaves = append(newLeaves, v)
				}
				break
			}
		}
		leaves = newLeaves
	}

	return leaves
}
