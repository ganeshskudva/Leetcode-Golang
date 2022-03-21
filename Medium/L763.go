package Medium

func partitionLabels(s string) []int {
	var res []int
	if len(s) == 0 {
		return res
	}

	mp := make(map[byte]int)
	right, size := 0, 0

	for i := 0; i < len(s); i++ {
		mp[s[i]] = i
	}

	for left := 0; left < len(s); left++ {
		size++
		if mp[s[left]] > right {
			right = mp[s[left]]
		}

		if left == right {
			res = append(res, size)
			size = 0
		}
	}

	return res
}
