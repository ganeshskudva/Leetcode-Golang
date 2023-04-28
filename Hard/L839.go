package Hard

func numSimilarGroups(strs []string) int {
	n := len(strs)
	parent := make([]int, n)
	var (
		isSimilar func(s1, s2 string) bool
		find      func(x int) int
	)

	isSimilar = func(s1, s2 string) bool {
		diff := 0
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				diff++
			}
		}

		return diff == 2 || diff == 0
	}

	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	for i := 0; i < n; i++ {
		parent[i] = i
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) {
				p1, p2 := find(i), find(j)
				if p1 != p2 {
					parent[p2] = p1
				}
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			res++
		}
	}

	return res
}
