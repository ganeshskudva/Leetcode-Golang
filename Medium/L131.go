func partition(s string) [][]string {
	var res [][]string
	dfs(s, 0, make([]string, 0), &res)

	return res
}

func dfs(s string, pos int, list []string, res *[][]string) {
	if pos == len(s) {
		tmp := make([]string, len(list))
		copy(tmp, list)
		*res = append(*res, tmp)
	} else {
		for i := pos; i < len(s); i++ {
			if isPalindrome(s, pos, i) {
				list = append(list, s[pos:i+1])
				dfs(s, i+1, list, res)
				list = list[:len(list)-1]
			}
		}
	}
}

func isPalindrome(s string, low, high int) bool {
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low, high = low+1, high-1
	}

	return true
}
