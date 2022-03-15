package Medium

type Stack struct {
	arr []int
}

func NewStack() *Stack {
	return &Stack{arr: make([]int, 0)}
}

func (s *Stack) Push(c int) {
	s.arr = append(s.arr, c)
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	ret := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return ret
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return 0
	}

	return s.arr[len(s.arr)-1]
}

func minRemoveToMakeValid(s string) string {
	var sb strings.Builder
	st := NewStack()
	match := make([]bool, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st.Push(i)
		} else if s[i] == ')' {
			if !st.IsEmpty() {
				// match these pairs, all unmatched are false anyway
				match[i] = true
				match[st.Pop()] = true
			}
		} else {
			// any character other than ( and ) are true anyway
			match[i] = true
		}
	}

	for i := 0; i < len(s); i++ {
		if match[i] {
			sb.WriteByte(s[i])
		}
	}

	if sb.Len() == 0 {
		return ""
	}

	return sb.String()
}
