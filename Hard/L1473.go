package Hard

const Max = math.MaxInt32

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := make(map[string]int)
	res := paint(houses, cost, n, 0, target, -1, dp)

	if res == Max {
		return -1
	} else {
		return res
	}
}

func paint(houses []int, cost [][]int, colors, houseIndex, remaining, lastColor int, dp map[string]int) int {
	key := strconv.Itoa(houseIndex) + "#" + strconv.Itoa(remaining) + "#" + strconv.Itoa(lastColor)

	if val, ok := dp[key]; ok {
		return val
	}

	if houseIndex == len(houses) && remaining == 0 {
		return 0
	}
	if houseIndex == len(houses) || remaining == -1 {
		return Max
	}

	minCost := Max
	if houses[houseIndex] == 0 {
		for color := 1; color <= colors; color++ {
			if color != lastColor {
				minCost = min(minCost, paint(houses, cost, colors, houseIndex+1, remaining-1, color, dp)+cost[houseIndex][color-1])
			} else {
				minCost = min(minCost, paint(houses, cost, colors, houseIndex+1, remaining, color, dp)+cost[houseIndex][color-1])
			}
		}
	} else {
		if houses[houseIndex] != lastColor {
			minCost = paint(houses, cost, colors, houseIndex+1, remaining-1, houses[houseIndex], dp)
		} else {
			minCost = paint(houses, cost, colors, houseIndex+1, remaining, lastColor, dp)
		}
	}
	dp[key] = minCost

	return dp[key]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
