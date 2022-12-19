package Easy

var exists struct{}

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph, vis := make([][]int, n), make(map[int]struct{})
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	
	return dfs(graph, vis, source, destination)
}

func dfs(graph [][]int, vis map[int]struct{}, source, dest int) bool {
	if source == dest {
		return true
	}
	if _, ok := vis[source]; ok {
		return false
	}
	vis[source] = exists
	
	found := false
	for _, neigh := range graph[source] {
		found = found || dfs(graph, vis, neigh, dest)
	}
	
	return found
}
