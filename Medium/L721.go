package Medium

func accountsMerge(accounts [][]string) [][]string {
	var ans [][]string
	if len(accounts) == 0 {
		return ans
	}

	owner, graph := make(map[string]string), make(map[string]map[string]bool)

	for _, a := range accounts {
		own := a[0]
		for i := 1; i < len(a); i++ {
			owner[a[i]] = own
			if _, ok := graph[a[i]]; !ok {
				graph[a[i]] = make(map[string]bool)
			}

			if i == 1 {
				continue
			}
			graph[a[i]][a[i-1]] = true
			graph[a[i-1]][a[i]] = true
		}
	}

	vis := make(map[string]bool)
	for k, _ := range graph {
		if _, ok := vis[k]; !ok {
			var tmp []string
			dfs(graph, k, &tmp, vis)
			sort.Strings(tmp)
			tmp = append([]string{owner[k]}, tmp...)
			ans = append(ans, tmp)
		}
	}

	return ans
}

func dfs(graph map[string]map[string]bool, s string, tmp *[]string, vis map[string]bool) {
	if _, ok := vis[s]; ok {
		return
	}

	vis[s] = true
	*tmp = append(*tmp, s)
	neighbors := graph[s]
	for nei, _ := range neighbors {
		dfs(graph, nei, tmp, vis)
	}
}
