package Medium

var exists struct{}

func possibleBipartition(n int, dislikes [][]int) bool {
	if n == 0 || len(dislikes) == 0 {
		return true
	}

	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for _, d := range dislikes {
		graph[d[0]] = append(graph[d[0]], d[1])
		graph[d[1]] = append(graph[d[1]], d[0])
	}

	color := make([]int, n+1)
	var q []int
	for i := 0; i <= n; i++ {
		if color[i] != 0 {
			continue
		}
		q = append(q, i)
		color[i] = 1
		st := make(map[int]struct{})

		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			if _, ok := st[node]; ok {
				continue
			}
			st[node] = exists

			for _, nei := range graph[node] {
				if color[nei] != 0 && color[nei] != -color[node] {
					return false
				}
				color[nei] = -color[node]
				q = append(q, nei)
			}
		}
	}

	return true
}
