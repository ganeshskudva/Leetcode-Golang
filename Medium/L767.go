package Medium

type Pair struct {
	Key   byte
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func reorganizeString(s string) string {
	freqMap, res := make(map[byte]int), make([]byte, len(s))
	for _, ch := range s {
		freqMap[byte(ch)]++
	}

	getSortedChars := func() PairList {
		p := make(PairList, len(freqMap))
		i := 0
		for k, v := range freqMap {
			p[i] = Pair{
				Key:   k,
				Value: v,
			}
			i++
		}
		sort.Sort(p)

		return p
	}

	sortedChars := getSortedChars()
	if freqMap[sortedChars[0].Key] > ((len(s) + 1) / 2) {
		return ""
	}

	idx := 0
	for _, pair := range sortedChars {
		for j := 0; j < freqMap[pair.Key]; j++ {
			if idx >= len(s) {
				idx = 1
			}
			res[idx] = pair.Key
			idx += 2
		}
	}
	return string(res)
}
