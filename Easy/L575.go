package Easy

import "fmt"

func distributeCandies(candyType []int) int {
	if len(candyType) == 0 {
		return 0
	}

	n := len(candyType)/2
	mp := make(map[int]bool)
	for _, v := range candyType {
		mp[v] = true
	}

	return min(len(mp), n)
}

func min(x, y int) int{
	if x < y {
		return x
	}

	return y
}
