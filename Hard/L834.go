package Hard

func sumOfDistancesInTree(n int, edges [][]int) []int {
	var (
		dfs1 func(node, parent int)
		dfs2 func(node, parent, dpVal int)
	)
	adjList, sz, val, ans := make([][]int, n), make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 0)
	}

	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	dfs1 = func(node, parent int) {
		for _, c := range adjList[node] {
			if c != parent {
				dfs1(c, node)
				sz[node] += sz[c]
				val[node] += val[c] + sz[c]
			}
		}
		sz[node]++
	}

	dfs2 = func(node, parent, dpVal int) {
		ans[node] = val[node] + dpVal + (n - sz[node])

		for _, c := range adjList[node] {
			if c != parent {
				dfs2(c, node, ans[node]-val[c]-sz[c])
			}
		}
	}

	dfs1(0, 0)
	dfs2(0, 0, 0)

	return ans
}
