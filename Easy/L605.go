package Easy

func canPlaceFlowers(flowerbed []int, n int) bool {
    cnt, res, size := 0, 0, len(flowerbed)

	for i := 0; i <= size; i++ {
		if i < size && flowerbed[i] == 0 {
			cnt++
			if i == 0 {
				cnt++
			}
			if i == size - 1 {
				cnt++
			}
		} else {
			res += (cnt - 1) / 2
			if res >= n {
				return true
			}
			cnt = 0
		}
	}

	return false
}
