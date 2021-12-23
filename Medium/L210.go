package Medium

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{}
	}

	ret, lst, indDegree, vis, cnt, idx := make([]int, numCourses), make([][]int, numCourses), make([]int, numCourses), make([]bool, numCourses), 0, 0
	if len(prerequisites) == 0 {
		for i := 0; i < numCourses; i++ {
			ret[i] = i
		}
		return ret
	}

	for i := 0; i < numCourses; i++ {
		lst[i] = make([]int, 0)
	}
	for _, p := range prerequisites {
		lst[p[1]] = append(lst[p[1]], p[0])
		indDegree[p[0]]++
	}

	for i := 0; i < numCourses; i++ {
		if indDegree[i] == 0 && !vis[i] {
			dfs(lst, ret, indDegree, vis, i, &idx, &cnt)
		}
	}

	if numCourses == cnt {
		return ret
	}

	return []int{}
}

func dfs(lst [][]int, ret, inDegree []int, vis []bool, c int, idx, cnt *int) {
	if vis[c] {
		return
	}

	vis[c] = true
	if inDegree[c] == 0 {
		ret[(*idx)] = c
		*idx++
		*cnt++
	}

	for _, nei := range lst[c] {
		inDegree[nei]--
		if inDegree[nei] == 0 {
			dfs(lst, ret, inDegree, vis, nei, idx, cnt)
		}
	}
}
