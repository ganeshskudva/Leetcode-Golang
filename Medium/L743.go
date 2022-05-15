package Medium

func networkDelayTime(times [][]int, n int, k int) int {
	dist := make([]int, n+1)
	for i := 1; i < len(dist); i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	for i := 1; i < n; i++ {
		for _, t := range times {
			source, dest, weight := t[0], t[1], t[2]
			if dist[source] != math.MaxInt32 && dist[source]+weight < dist[dest] {
				dist[dest] = dist[source] + weight
			}
		}
	}

	max := math.MinInt32
	for i := range dist {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		if dist[i] > max {
			max = dist[i]
		}
	}

	return max
}
