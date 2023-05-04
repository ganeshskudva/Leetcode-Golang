package Medium

func predictPartyVictory(senate string) string {
	var q1, q2 []int
	n := len(senate)
	for i := 0; i < n; i++ {
		if senate[i] == 'R' {
			q1 = append(q1, i)
		} else {
			q2 = append(q2, i)
		}
	}
	
	for len(q1) > 0 && len(q2) > 0 {
		rIndex, dIndex := q1[0], q2[0]
		q1, q2 = q1[1:], q2[1:]
		
		if rIndex < dIndex {
			q1 = append(q1, rIndex + n)
		} else {
			q2 = append(q2, dIndex + n)
		}
 	}
	if len(q1) > len(q2) {
		return "Radiant"
	} 
	
	return "Dire"
}
