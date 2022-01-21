package Medium

func canCompleteCircuit(gas []int, cost []int) int {
	sumGas, sumCost, start, tank := 0, 0, 0, 0
	for i := range gas {
		sumGas += gas[i]
		sumCost += cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			start, tank = i+1, 0
		}
	}

	if sumGas < sumCost {
		return -1
	}

	return start
}
