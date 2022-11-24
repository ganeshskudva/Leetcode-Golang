package Hard

func minMovesToMakePalindrome(s string) int {
	l, r, k, cnt := 0, len(s)-1, len(s)-1, 0
	arr := []byte(s)

	for l < r {
		if arr[l] == arr[r] {
			l, r = l+1, r-1
		} else {
			k = findKthIndexMatchingwithLthIndex(arr, l, r)

			if k == l {
				if l+1 < len(s) {
					arr[l], arr[l+1] = arr[l+1], arr[l]
				}
				cnt++
			} else {
				for k < r {
					if k+1 < len(s) {
						arr[k], arr[k+1] = arr[k+1], arr[k]
					}
					cnt++
					k++
				}
				l, r = l+1, r-1
			}
		}
	}

	return cnt
}

func findKthIndexMatchingwithLthIndex(arr []byte, l, k int) int {
	for k > l {
		if arr[k] == arr[l] {
			return k
		}
		k--
	}
	return k
}
