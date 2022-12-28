package Medium

func probabilityOfHeads(prob []float64, target int) float64 {
	mp := make(map[string]float64)

	return solve(len(prob), target, prob, mp)
}

func solve(idx, tgt int, prob []float64, mp map[string]float64) float64 {
	if idx == 0 && tgt == 0 {
		return 1
	}
	if idx == 0 {
		return 0
	}
	if idx < tgt {
		return 0
	}

	key := fmt.Sprintf("%d-%d", idx, tgt)
	if v, ok := mp[key]; ok {
		return v
	}

	if tgt == 1 {
		mp[key] += solve(idx-1, tgt-1, prob, mp) * prob[idx-1]
	}
	mp[key] += solve(idx-1, tgt, prob, mp) * (1 - prob[idx-1])

	return mp[key]
}
