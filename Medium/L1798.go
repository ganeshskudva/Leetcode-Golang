package Medium

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	cnt := 1
	for i:=0; i< len(coins) && coins[i] <= cnt; i++ {
		cnt += coins[i]
	}
	
	return cnt
}
