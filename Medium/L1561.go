package Medium

func maxCoins(piles []int) int {
	sort.Ints(piles)
	n, maximum := len(piles), 0

	for i := n / 3; i < n-1; i += 2 {
		maximum += piles[i]
	}

	return maximum
}
