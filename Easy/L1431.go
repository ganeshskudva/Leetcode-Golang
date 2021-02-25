package Easy

import "math"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := math.MinInt32
	var ans = make([]bool, len(candies))
	if len(candies) == 0 {
		return nil
	}

	for _, val := range candies {
		max = maxi(max, val)
	}

	for i, val:= range candies {
		if val + extraCandies >= max {
			ans[i] = true
		}
	}

	return ans
}

func maxi(a, b int) int {
	if a > b {
		return a
	}

	return b
}
