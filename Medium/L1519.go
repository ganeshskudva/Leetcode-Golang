package Medium

func countSubTrees(n int, edges [][]int, labels string) []int {
	res, mp := make([]int, n), make(map[int][]int)

	for _, e := range edges {
		mp[e[0]] = append(mp[e[0]], e[1])
		mp[e[1]] = append(mp[e[1]], e[0])
	}

	solve(mp, labels, 0, -1, &res)
	return res
}

func solve(graph map[int][]int, labels string, node, parent int, res *[]int) []int {
	nodeCnt := make([]int, 26)
	nodeCnt[labels[node]-'a']++

	for _, nei := range graph[node] {
		if nei != parent {
			childCnt := solve(graph, labels, nei, node, res)
			for i := 0; i < 26; i++ {
				nodeCnt[i] += childCnt[i]
			}
		}
	}

	(*res)[node] = nodeCnt[labels[node]-'a']
	return nodeCnt
}
