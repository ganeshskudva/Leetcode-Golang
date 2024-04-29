package Medium

func minOperations(nums []int, k int) int {
	finalXor := 0
	for _, n := range nums {
		finalXor ^= n
	}
	
	cnt := 0
	for k > 0 || finalXor > 0 {
		if (k % 2) != (finalXor % 2) {
			cnt++
		}

		k /= 2
		finalXor /= 2
	}
	
	return cnt
}
