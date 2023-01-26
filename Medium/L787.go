package Medium

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[src] = 0

	for i := 0; i <= k; i++ {
		tmp := make([]int, len(cost))
		copy(tmp, cost)
		for _, f := range flights {
			curr, nxt, price := f[0], f[1], f[2]
			if cost[curr] == math.MaxInt32 {
				continue
			}
			tmp[nxt] = min(tmp[nxt], cost[curr]+price)
		}
		cost = tmp
	}

	if cost[dst] == math.MaxInt32 {
		return -1
	}

	return cost[dst]
}
