package Medium

type pair struct {
	x, y int
}

func newPair(x, y int) pair {
	return pair{
		x: x,
		y: y,
	}
}

func minScore(n int, roads [][]int) int {
	adj, vis := make([][]pair, n+1), make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		adj[i] = make([]pair, 0)
	}

	for i := range roads {
		adj[roads[i][0]] = append(adj[roads[i][0]], newPair(roads[i][1], roads[i][2]))
		adj[roads[i][1]] = append(adj[roads[i][1]], newPair(roads[i][0], roads[i][2]))
	}

	var q []int
	q = append(q, 1)
	vis[1] = true
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, p := range adj[node] {
			if !vis[p.x] {
				q = append(q, p.x)
				vis[p.x] = true
			}
		}
	}

	minPath := math.MaxInt32
	for i := range roads {
		if vis[roads[i][0]] && vis[roads[i][1]] && roads[i][2] < minPath {
			minPath = roads[i][2]
		}
	}

	return minPath
}
