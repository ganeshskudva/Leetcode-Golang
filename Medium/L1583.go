package Medium

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
    graph := make([]int, n)

	for _, p := range pairs {
		graph[p[0]] = p[1]
		graph[p[1]] = p[0]
	}

	res := 0
	type M map[int]int
	prefer := make([]M, n)

	for i := 0; i < n; i++ {
		for j:= 0; j < n-1; j++ {
			if prefer[i] == nil {
				prefer[i] = make(M)
			}
			prefer[i][preferences[i][j]] = j
		}
	}

	for i:= 0; i < n; i++ {
		for _, j := range preferences[i] {
			if prefer[j][i] < prefer[j][graph[j]] && 
				prefer[i][graph[i]] > prefer[i][j]{
				res++
				break
			}
		}
	}
	
	return res
}
