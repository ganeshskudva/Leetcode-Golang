func threeEqualParts(arr []int) []int {
	numOne := 0
	for _, i := range arr {
		if i == 1 {
			numOne++
		}
	}

	noRes := []int{-1, -1}
	if numOne == 0 {
		return []int{0, 2}
	}
	if numOne%3 != 0 {
		return noRes
	}

	idxThird, temp := 0, 0
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 1 {
			temp++
			if temp == numOne/3 {
				idxThird = i
				break
			}
		}
	}

	res1 := findEndIdx(arr, 0, idxThird)
	if res1 < 0 {
		return noRes
	}

	res2 := findEndIdx(arr, res1+1, idxThird)
	if res2 < 0 {
		return noRes
	}

	return []int{res1, res2 + 1}
}

func findEndIdx(A []int, left, right int) int {
	for A[left] == 0 {
		left++
	}
	for right < len(A) {
		if A[left] != A[right] {
			return -1
		}
		left++
		right++
	}

	return left - 1
}
