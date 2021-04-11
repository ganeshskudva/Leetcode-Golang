package Medium

func shoppingOffers(price []int, special [][]int, needs []int) int {
	return bt(special, price, &needs, 0)
}

func bt(special [][]int, price []int, needs *[]int, pos int) int {
	min := directPurchase(price, *needs)

	for i := pos; i < len(special); i++ {
		offer := special[i]
		var tmp []int
		for j := 0; j < len(*needs); j++ {
			if (*needs)[j] < offer[j] {
				tmp = nil
				break
			}
			tmp = append(tmp, (*needs)[j]-offer[j])
		}

		if len(tmp) != 0 {
			min = mini(min, offer[len(offer)-1]+bt(special, price, &tmp, i))
		}
	}

	return min
}

func directPurchase(price, needs []int) int {
	total := 0
	for i := 0; i < len(needs); i++ {
		total += price[i] * needs[i]
	}

	return total
}

func mini(a, b int) int {
	if a < b {
		return a
	}

	return b
}
