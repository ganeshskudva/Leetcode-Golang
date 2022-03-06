package Hard

func countOrders(n int) int {
	dp := make([][]int64, 501)
	for i := 0; i < 501; i++ {
		dp[i] = make([]int64, 501)
		for j := 0; j < 501; j++ {
			dp[i][j] = -1
		}
	}

	return int(solve(n, n, &dp))
}

func solve(pick, del int, dp *[][]int64) int64 {
	if del < 0 || pick < 0 {
		//base case as it can't be negative
		return 0
	}

	if del < pick {
		// condition to insure we are first picking then delivering
		return 0
	}

	if pick == 0 && del == 0 {
		//everything finishes return 1
		return 1
	}

	if (*dp)[pick][del] != -1 {
		//memoization
		return (*dp)[pick][del]
	}

	var ans int64
	mod := 1000000007

	// if there are pick number of orders we have [pick]C1 = pick number of choices for selecting  and we have picked one so we do pick-1
	ans += int64(pick) * solve(pick-1, del, dp)
	ans %= int64(mod)

	// delivering : the number of items we have picked but not delivered is (del-pick) so we can select one item from these to deliver [del-pick]C1  =del-pick and we have delivered so del-1
	ans += int64(del-pick) * solve(pick, del-1, dp)
	ans %= int64(mod)

	(*dp)[pick][del] = ans

	return (*dp)[pick][del]
}
