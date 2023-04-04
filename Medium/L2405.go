package Medium

func partitionString(s string) int {
  var exists struct{}
	idx, cnt := 0, 0
	mp := make(map[byte]struct{})

	for i := range s {
		if _, ok := mp[s[i]]; ok {
			cnt++
			mp = make(map[byte]struct{})
		}
		mp[s[i]] = exists
		idx++
	}
	
	return cnt + 1
}
