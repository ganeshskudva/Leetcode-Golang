package Medium

func canFinish(numCourses int, prerequisites [][]int) bool {
	var exists struct{}
	if numCourses == 0 {
		return false
	}

	inDegree, mp := make([]int, numCourses), make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		mp[i] = make([]int, 0)
	}

	for _, p := range prerequisites {
		mp[p[1]] = append(mp[p[1]], p[0])
		inDegree[p[0]]++
	}

	var q []int
	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}
	st := make(map[int]struct{})
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			c := q[0]
			q = q[1:]

			if _, ok := st[c]; ok {
				continue
			}
			st[c] = exists
			for _, n := range mp[c] {
				inDegree[n]--
				if inDegree[n] == 0 {
					q = append(q, n)
				}
			}
		}
	}

	return len(st) == numCourses
}
