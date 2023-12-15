package Easy

func destCity(paths [][]string) string {
	inDeg, outDeg := make(map[string]int), make(map[string]int)

	for _, p := range paths {
		outDeg[p[0]] += 1
		inDeg[p[1]] += 1
	}

	for k, v := range inDeg {
		if v == 1 && outDeg[k] == 0 {
			return k
		}
	}

	return ""
}
