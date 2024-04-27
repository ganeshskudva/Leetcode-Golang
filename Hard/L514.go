package Hard

func findRotateSteps(ring string, key string) int {
	var solve func(keyIdx, ringIdx int) int
	mp := make(map[string]int)

	solve = func(keyIdx, ringIdx int) int {
		if keyIdx == len(key) {
			return 0
		}

		mpKey := fmt.Sprintf("%s%d", keyIdx, ringIdx)
		if v, ok := mp[mpKey]; ok {
			return v
		}

		size, stepsReq, tmp, cnt := len(ring), 1000000000, ringIdx, 0

		for cnt < size {
			if key[keyIdx] == ring[tmp] {
				stepsReq = min(stepsReq, 1+cnt+solve(keyIdx+1, tmp))
			}
			cnt, tmp = cnt+1, tmp+1
			
			tmp %= size
		}

		cnt, tmp = 0, ringIdx
		for cnt < size {
			if key[keyIdx] == ring[tmp] {
				stepsReq = min(stepsReq, 1+cnt+solve(keyIdx+1, tmp))
			}
			cnt, tmp = cnt+1, tmp-1
			tmp += size
			tmp %= size
		}

		mp[mpKey] = stepsReq
		return stepsReq
	}

	return solve(0, 0)
}
