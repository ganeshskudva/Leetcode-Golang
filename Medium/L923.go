package Medium

func threeSumMulti(arr []int, target int) int {
	c := make([]float64, 101)

	for _, a := range arr {
		c[a]++
	}
	res := 0.0

	for i := 0; i <= 100; i++ {
		for j := i; j <= 100; j++ {
			k := target - i - j
			if k > 100 || k < 0 {
				continue
			}
			if i == j && j == k {
				res += c[i] * (c[i] - 1) * (c[i] - 2) / 6
			} else if i == j && j != k {
				res += c[i] * (c[i] - 1) / 2 * c[k]
			} else if j < k {
				res += c[i] * c[j] * c[k]
			}
		}
	}

	return int(res) % (1e9 + 7);
}
