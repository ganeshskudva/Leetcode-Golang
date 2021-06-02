package Medium

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	cache := make(map[int]bool)
	return solve(s1, s2, s3, 0, 0, cache)
}

func solve(s1, s2, s3 string, p1, p2 int, cache map[int]bool) bool{
	if p1 + p2 == len(s3) {
		return true
	}
	
	if _, ok := cache[(p1*len(s3))+p2]; ok {
		return false
	}
	
	cache[(p1*len(s3))+p2] = true
	var match1, match2 bool
	if p1 < len(s1) && s3[p1+p2] == s1[p1] {
		match1 = true
	}
	if p2 < len(s2) && s3[p1+p2] == s2[p2] {
		match2 = true
	}
	
	if match1 && match2 {
		return solve(s1, s2, s3, p1+1, p2, cache ) || solve(s1, s2, s3, p1, p2+1, cache)
	} else if match1 {
		return solve(s1, s2, s3, p1+1, p2, cache)
	} else if match2 {
		return solve(s1, s2, s3, p1, p2+1, cache)
	} else {
		return false
	}
}
