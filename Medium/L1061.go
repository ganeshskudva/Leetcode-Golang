package Medium

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	var sb strings.Builder
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = -1
	}

	for i := 0; i < len(s1); i++ {
		union(int(s1[i]-'a'), int(s2[i]-'a'), &parent)
	}
	for i := 0; i < len(baseStr); i++ {
		sb.WriteByte(byte(find(int(baseStr[i]-'a'), &parent) + 'a'))
	}

	return sb.String()
}

func union(x, y int, parent *[]int) {
	x, y = find(x, parent), find(y, parent)

	if x != y {
		(*parent)[max(x, y)] = min(x, y)
	}
}

func find(x int, parent *[]int) int {
	if (*parent)[x] == -1 {
		return x
	}

	(*parent)[x] = find((*parent)[x], parent)
	return (*parent)[x]
}
