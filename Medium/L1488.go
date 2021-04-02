package Medium

func avoidFlood(rains []int) []int {
	avoid := make([]int, len(rains))
	last := map[int]int{}
	valid := make([]int, 0)
	for i, v := range rains {
		if v > 0 {
			avoid[i] = -1
			if _, exist := last[v]; exist {
				if index := binarySearch(valid, last[v]); index < len(valid) {
					avoid[valid[index]] = v
					copy(valid[index:], valid[index+1:])
					valid = valid[:len(valid)-1]
				} else {
					return []int{}
				}
			}
			last[v] = i
		} else {
			valid = append(valid, i)
		}
	}
	for i, v := range avoid {
		if v == 0 {
			avoid[i] = 1
		}
	}
	return avoid
}

func binarySearch(array []int, value int) int {
	l, r := 0, len(array)-1
	for l <= r {
		m := (l+r)>>1
		if array[m] < value {
			l = m+1
		} else {
			r = m-1
		}
	}
	return l
}
