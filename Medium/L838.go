package Medium

func pushDominoes(dominoes string) string {
	arr := []byte(dominoes)
	L, R := -1, -1
	for i := 0; i <= len(dominoes); i++ {
		if i == len(arr) || arr[i] == 'R' {
			if R > L {
				for R < i {
					arr[R] = 'R'
					R++
				}
			}
			R = i
		} else if arr[i] == 'L' {
			if L > R || R == -1 && L == -1 {
				L++
				for L < i {
					arr[L] = 'L'
					L++
				}
			} else {
				L = i
				lo, hi := R+1, L-1
				for lo < hi {
					arr[lo] = 'R'
					arr[hi] = 'L'
					lo++
					hi--
				}
			}
		}
	}

	return string(arr)
}
