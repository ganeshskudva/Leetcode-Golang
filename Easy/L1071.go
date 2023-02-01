package Easy

func gcdOfStrings(str1 string, str2 string) string {
	s1, s2 := fmt.Sprintf("%s%s", str1, str2), fmt.Sprintf("%s%s", str2, str1)
	if s1 != s2 {
		return ""
	}

	gcdVal := gcd(len(str1), len(str2))
	return str2[:gcdVal]
}

func gcd(p, q int) int {
	if q == 0 {
		return p
	}

	return gcd(q, p%q)
}
