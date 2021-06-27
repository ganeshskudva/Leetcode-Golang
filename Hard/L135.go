package Hard

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}

	for i := range candies {
		if i == 0 {
			continue
		}

		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(candies) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i+1]+1 > candies[i] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	sum := 0
	for _, c := range candies {
		sum += c
	}

	return sum
}
