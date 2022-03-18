package Easy

func canBeEqual(target []int, arr []int) bool {
	m, n := len(arr), len(target)
	if m != n {
		return false
	}

	mp := make(map[int]int)
	for i := 0; i < m; i++ {
		mp[target[i]]++
		mp[arr[i]]--
	}

	for _, v := range mp {
		if v != 0 {
			return false
		}
	}

	return true
}
