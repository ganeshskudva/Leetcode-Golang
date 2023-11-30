package Medium

func minimumOneBitOperations(n int) int {
	var solve func(n, res int) int
	
	solve = func(n, res int) int {
		if n == 0 {
			return res
		}
		b := 1
		for (b << 1) <= n {
			b = b << 1
		}
		
		return solve((b >> 1) ^ b ^ n, res + b)
	}
	
	return solve(n, 0)
}
