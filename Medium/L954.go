package Medium

func canReorderDoubled(arr []int) bool {
	if len(arr) == 0 {
		return true
	}
	sort.Ints(arr)
	mp := make(map[int]int)
	for _, i := range arr {
		if i == 0 {
			continue
		}
		mp[i]++
	}

	for i := range arr {
		if arr[i] != 0 {
			if mp[arr[i]] > 0 {
				tgt := arr[i] * 2
				if arr[i] < 0 {
					if arr[i]%2 != 0 {
						return false
					} else {
						tgt = arr[i] / 2
					}
				}
				val, exists := mp[tgt]
				if !exists || val < mp[arr[i]] {
					return false
				} else {
					mp[tgt] = mp[tgt] - mp[arr[i]]
					mp[arr[i]] = 0
				}
			}
		}
	}

	return true
}
