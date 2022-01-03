package Easy

var exists struct{}

func findJudge(n int, trust [][]int) int {
	mp, inDegree := make(map[int]struct{}), make([]int, n+1)

	for i := 1; i <= n; i++ {
		mp[i] = exists
	}

	for _, t := range trust {
		delete(mp, t[0])
		inDegree[t[1]]++
	}

	if len(mp) == 1 {
		for k, _ := range mp {
			if inDegree[k] == n-1 {
				return k
			}
		}
	}

	return -1
}
