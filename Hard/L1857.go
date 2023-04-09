package Hard

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func largestPathValue(colors string, edges [][]int) int {
	var graph [][]int
	n, inDegree := len(colors), make([]int, len(colors))

	for i := 0; i < n; i++ {
		graph = append(graph, make([]int, 0))
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		inDegree[v]++
		graph[u] = append(graph[u], v)
	}

	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, 26)
	}

	var q []int
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
			mp[i][colors[i]-'a'] = 1
		}
	}

	getMax := func(nums []int) int {
		maxi := math.MinInt32

		for _, n := range nums {
			maxi = max(maxi, n)
		}

		return maxi
	}

	var res, seen int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		seen++

		res = max(res, getMax(mp[node]))

		for _, nei := range graph[node] {
			for i := 0; i < 26; i++ {
				val := mp[node][i]
				if int(colors[nei]-'a') == i {
					val++
				}
				mp[nei][i] = max(mp[nei][i], val)
			}
			inDegree[nei]--

			if inDegree[nei] == 0 {
				q = append(q, nei)
			}
		}
	}

	if seen == n {
		return res
	}
	
	return -1 
}
