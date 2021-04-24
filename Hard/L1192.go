package Hard

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func criticalConnections(n int, connections [][]int) [][]int {
	var res [][]int
	var timer int
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = []int{}
	}

	for _, c := range connections {
		graph[c[0]] = append(graph[c[0]], c[1])
		graph[c[1]] = append(graph[c[1]], c[0])
	}
	visited, timeStampAtThatNode, timer := make([]bool, n), make([]int, n), 0
	dfs(graph, -1, 0, &timer, &visited, &res, &timeStampAtThatNode)
	return res
}

func dfs(graph [][]int, parent, node int, timer *int, visited *[]bool, res *[][]int, timeStampAtThatNode *[]int) {
	(*visited)[node] = true
	*timer++
	(*timeStampAtThatNode)[node] = *timer
	currentTimeStamp := (*timeStampAtThatNode)[node]

	for _, neigh := range graph[node] {
		if neigh == parent {
			continue
		}
		if !(*visited)[neigh] {
			dfs(graph, node, neigh, timer, visited, res, timeStampAtThatNode)
		}
		(*timeStampAtThatNode)[node] = min((*timeStampAtThatNode)[node], (*timeStampAtThatNode)[neigh])
		if currentTimeStamp < (*timeStampAtThatNode)[neigh] {
			*res = append(*res, []int{node, neigh})
		}
	}
}
