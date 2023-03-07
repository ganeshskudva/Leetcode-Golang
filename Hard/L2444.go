package Hard

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var (
		res, start int64
		minStart, maxStart int64
	)
	minFound, maxFound := false, false
	
	for i := range nums {
		n := nums[i]
		if n < minK || n > maxK {
			minFound, maxFound, start = false, false, int64(i+1)
		}
		if n == minK {
			minFound, minStart = true, int64(i)
		}
		if n == maxK {
			maxFound, maxStart = true, int64(i)
		}
		if minFound && maxFound {
			res += min(minStart, maxStart) - start + 1
		}
	}
	
	return res
}

func min(a, b int64) int64 {
  if a < b {
    return a
  }

  return b
}
