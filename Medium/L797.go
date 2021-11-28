package Medium

func allPathsSourceTarget(graph [][]int) [][]int {
	var ret, lst [][]int
	if len(graph) == 0 {
		return ret
	}

	for i := range graph {
		lst = append(lst, make([]int, 0))
		for _, g := range graph[i] {
			lst[i] = append(lst[i], g)
		}
	}

	bt(&ret, lst, make([]int, 0), 0, len(graph)-1)
	return ret
}

func bt(ret *[][]int, lst [][]int, tmp []int, index, n int) {
	if index == n {
		tmp = append(tmp, index)
		temp := make([]int, len(tmp))
		copy(temp, tmp)
		*ret = append(*ret, temp)
		tmp = tmp[:len(tmp)-1]
		return
	}

	tmp = append(tmp, index)
	for _, nei := range lst[index] {
		bt(ret, lst, tmp, nei, n)
	}
	tmp = tmp[:len(tmp)-1]
}
