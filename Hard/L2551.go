package Hard

func putMarbles(weights []int, k int) int64 {
	n, pairs := len(weights), make([]int, len(weights)-1)

	for i := 1; i < n; i++ {
		pairs[i-1] = weights[i] + weights[i-1]
	}
	sort.Ints(pairs)
	minScore, maxScore := 0, 0
	for i := 0; i < k-1; i++ {
		minScore += pairs[i]
		maxScore += pairs[n-i-2]
	}

	return int64(maxScore - minScore)
}
