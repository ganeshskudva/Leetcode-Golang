package Easy

func removeDuplicates(s string) string {
	st := NewStack()

	for _, c := range s {
		if !st.isEmpty() {
			if st.peek() == byte(c) {
				st.pop()
				continue
			} else {
				st.push(byte(c))
			}
		} else {
			st.push(byte(c))
		}
	}

	var sb strings.Builder
	for !st.isEmpty() {
		sb.WriteByte(st.pop())
	}

	return reverse(sb.String())
}

type Stack struct {
	st []byte
}

func NewStack() *Stack {
	return &Stack{st: []byte{}}
}

func (s *Stack) pop() byte {
	ret := s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]

	return ret
}

func (s *Stack) peek() byte {
	return s.st[len(s.st)-1]
}

func (s *Stack) push(c byte) {
	s.st = append(s.st, c)
}

func (s *Stack) isEmpty() bool {
	return len(s.st) == 0
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
