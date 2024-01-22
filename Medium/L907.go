package Medium

func sumSubarrayMins(arr []int) int {
	var (
		res int64
		st []int
		mod = int64(1000000007)
        
	)
	st = append(st, -1)
	
	for i := 0; i < len(arr) + 1; i++ {
		currVal := 0
		if i < len(arr) {
			currVal = arr[i]
		}
		
		for st[len(st) - 1] != -1 && currVal < arr[st[len(st) - 1]] {
			index := st[len(st) - 1]
			st = st[:len(st) - 1]
			newTop := st[len(st) - 1]
			left, right := index - newTop, i - index
			add := int64(left * right * arr[index]) % mod
			res += add
			res %= mod
		}
		
		st = append(st, i)
	}
	
	return int(res)
}
