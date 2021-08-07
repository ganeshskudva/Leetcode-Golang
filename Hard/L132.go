package Hard

func minCut(s string) int {
	cut, n := make([]int, len(s)), len(s)
	pal := make([][]bool, n*n)
	for i := 0; i < n; i++ {
		pal[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		min := i
		for j := 0; j <= i; j++ {
			if s[j] == s[i] && (j+1 > i-1 || pal[j+1][i-1]) {
				pal[j][i] = true
				if j == 0 {
					min = 0
				} else {
					if cut[j-1]+1 < min {
						min = cut[j-1] + 1
					}
				}
			}
		}
		cut[i] = min
	}

	return cut[n-1]
}
