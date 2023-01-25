package Medium

var exists struct{}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n, ans, minDist := len(edges), -1, math.MaxInt32
	dist1, dist2 := make([]int, n), make([]int, n)
	vis1, vis2 := make(map[int]struct{}), make(map[int]struct{})

	solve(edges, node1, &dist1, vis1)
	solve(edges, node2, &dist2, vis2)

	for i := 0; i < n; i++ {
		if hasVisited(i, vis1) && hasVisited(i, vis2) {
			if minDist > max(dist1[i], dist2[i]) {
				minDist = max(dist1[i], dist2[i])
				ans = i
			}
		}
	}

	return ans
}

func hasVisited(node int, vis map[int]struct{}) bool {
	if _, ok := vis[node]; ok {
		return true
	}

	return false
}

func solve(edges []int, node int, dist *[]int, vis map[int]struct{}) {
	vis[node] = exists
	neigh := edges[node]
	if neigh != -1 {
		if !hasVisited(neigh, vis) {
			(*dist)[neigh] = (*dist)[node] + 1
			solve(edges, neigh, dist, vis)
		}
	}
}
