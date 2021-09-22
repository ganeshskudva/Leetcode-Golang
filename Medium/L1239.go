package Medium

func maxLength(arr []string) int {
	result := 0

	if len(arr) == 0 {
		return result
	}

	dfs(arr, "", 0, &result)

	return result
}

func dfs(arr []string, path string, idx int, result *int) {
	unique := isUnique(path)

	if unique {
		if len(path) > *result {
			*result = len(path)
		}
	}

	if idx == len(arr) || !unique {
		return
	}

	for i := idx; i < len(arr); i++ {
		dfs(arr, path+arr[i], i+1, result)
	}
}

func isUnique(s string) bool {
	mp := make(map[byte]bool)

	for _, c := range s {
		if _, ok := mp[byte(c)]; ok {
			return false
		}
		mp[byte(c)] = true
	}

	return true
}
