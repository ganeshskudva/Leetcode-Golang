package Medium

func equationsPossible(equations []string) bool {
	root := make([]int, 26)

	for i := 0; i < 26; i++ {
		root[i] = i
	}
	for _, e := range equations {
		if e[1] == '=' {
			root[find(int(e[0]-'a'), &root)] = find(int(e[3]-'a'), &root)
		}
	}

	for _, e := range equations {
		if e[1] == '!' && find(int(e[0]-'a'), &root) == find(int(e[3]-'a'), &root) {
			return false
		}
	}

	return true
}

func find(x int, root *[]int) int {
	if x != (*root)[x] {
		(*root)[x] = find((*root)[x], root)
	}

	return (*root)[x]
}
