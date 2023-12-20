package Easy

func buyChoco(prices []int, money int) int {
	if len(prices) < 2 {
		return money
	}
	
	first, second := math.MaxInt32, math.MaxInt32
	
	for _, p := range prices {
		if p < first {
			second, first = first, p
		}else if p < second {
			second = p
		}
	}
	
	if first + second > money {
		return money
	}
	
	return money - (first + second)
}
