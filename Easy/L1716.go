package Easy

func totalMoney(n int) int {
	sum, prev, tot, cnt := 0, 1, 0, 0

	for i := 1; i <= n; i++ {
		if cnt == 7 {
			cnt, tot = 0, prev+1
			prev = tot
		} else {
			tot++
		}

		sum += tot
		cnt++
	}

	return sum
}
