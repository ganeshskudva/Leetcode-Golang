package Easy

func toLowerCase(s string) string {
  var sb strings.Builder
	
	for i := range s {
		ch := s[i]
		if s[i] >='A' && s[i] <= 'Z' {
			ch = 'a' + (s[i] - 'A')
		}
		sb.WriteByte(ch)
	}
	
	return sb.String()
}
