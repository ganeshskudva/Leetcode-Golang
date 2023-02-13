package Medium

var exists struct{}

//1=red, 2=blue, 0=root-edge (special case)
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	reds, blues := make([][]int, n), make([][]int, n)

	for _, e := range redEdges {
		reds[e[0]] = append(reds[e[0]], e[1])
	}

	for _, e := range blueEdges {
		blues[e[0]] = append(blues[e[0]], e[1])
	}

	var q [][]int
	res, moves, seen := make([]int, n), 0, make(map[string]struct{})
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	q = append(q, []int{0, 0})

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			curr := q[0]
			q = q[1:]
			key := fmt.Sprintf("%d-%d", curr[0], curr[1])
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = exists

			if res[curr[0]] == -1 {
				res[curr[0]] = moves
			}
			if curr[1] == 2 || curr[1] == 0 {
				if len(reds[curr[0]]) > 0 {
					for _, neigh := range reds[curr[0]] {
						q = append(q, []int{neigh, 1})
					}
				}
			}
			if curr[1] == 1 || curr[1] == 0 {
				if len(blues[curr[0]]) > 0 {
					for _, neigh := range blues[curr[0]] {
						q = append(q, []int{neigh, 2})
					}
				}
			}
		}
		moves++
	}

	return res
}
