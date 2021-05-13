package Medium

func ambiguousCoordinates(s string) []string {
	var res []string
	s = s[1:len(s)-1]
	for i := 1; i < len(s); i++ {
		left, right := valid(s[0:i]), valid(s[i:])
		for _, l:= range left {
			for _, r := range right {
				res = append(res, "(" + l + ", " + r + ")")
			}
		}
	}
	
	return res
}

func valid(s string) []string {
	n := len(s)
	var res []string
	if s[0] == '0' && s[n-1] == '0' {
		if n == 1 {
			res = append(res, "0")
		}
		return res
	}
	
	if s[0] == '0' {
		res = append(res, "0."+s[1:])
		return res
	}
	if s[n-1] == '0' {
		res = append(res, s)
		return res
	}
	res = append(res, s)
	for i:= 1; i < n; i++ {
		res = append(res, s[0:i] + "." + s[i:])
	}
	
	return res
}
