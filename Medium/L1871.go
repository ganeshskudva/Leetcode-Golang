package Medium

func canReach(s string, minJump int, maxJump int) bool {
    	n, farthest := len(s), 0
	var q []int
	q = append(q, 0)
	
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		start := max(curr + minJump, farthest+1)
		for j:= start; j < curr + maxJump + 1; j++ {
			if j >= n {
				break
			}
			if s[j] == '0' {
				if j == n-1 {
					return true
				}
				q = append(q, j)
			}
		}
		farthest = curr + maxJump
	}
	
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
