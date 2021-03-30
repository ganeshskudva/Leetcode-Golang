package Hard

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		} else {
			return envelopes[i][0] < envelopes[j][0]
		}
	})

	var dp []int
	for _, envelope := range envelopes {
		i := sort.SearchInts(dp, envelope[1])
		if i == len(dp) {
			dp = append(dp, envelope[1])
		} else {
			dp[i] = envelope[1]
		}
	}

	return len(dp)
}
