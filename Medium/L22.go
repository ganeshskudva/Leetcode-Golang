package Medium

func generateParenthesis(n int) []string {
	var ans []string
	var tmp string
	if n == 0 {
		return ans
	}

	dfs(&ans, &tmp, n, 0, 0)

	return ans
}

func dfs(ans *[]string, str *string, n, open, close int) {
	if len(*str) == 2*n {
		*ans = append(*ans, *str)
		return
	}

	if open < n {
		*str += "("
		dfs(ans, str, n, open+1, close)
		*str = (*str)[:len(*str)-1]
	}

	if close < open {
		*str += ")"
		dfs(ans, str, n, open, close+1)
		*str = (*str)[:len(*str)-1]
	}
}
