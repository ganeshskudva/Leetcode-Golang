package Medium

type Deque struct {
	operands  []int
	operators []byte
}

func NewDeque() *Deque {
	return &Deque{
		operands:  make([]int, 0),
		operators: make([]byte, 0),
	}
}

func (d *Deque) PushOperand(val int) {
	d.operands = append(d.operands, val)
}

func (d *Deque) PopOperand() int {
	val := d.operands[len(d.operands)-1]
	d.operands = d.operands[:len(d.operands)-1]

	return val
}

func (d *Deque) PushOperator(ch byte) {
	d.operators = append(d.operators, ch)
}

func (d *Deque) PopOperator() byte {
	val := d.operators[len(d.operators)-1]
	d.operators = d.operators[:len(d.operators)-1]

	return val
}

func (d *Deque) PeekOperator() byte {
	return d.operators[len(d.operators)-1]
}

func (d *Deque) isOperatorEmpty() bool {
	return len(d.operators) == 0
}

func (d *Deque) operate() int {
	a, b, c := d.PopOperand(), d.PopOperand(), d.PopOperator()

	switch c {
	case '+':
		return a + b
	case '-':
		return b - a
	case '*':
		return a * b
	case '/':
		return b / a
	default:
		return 0
	}
}

var precedenceMap = map[byte]int{
	'(': -1,
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
}

func calculate(s string) int {
	q, n := NewDeque(), len(s)

	for i := 0; i < n; i++ {
		c := s[i]
		if isDigit(c) {
			val := int(c - '0')
			for i+1 < n && isDigit(s[i+1]) {
				val = val*10 + int(s[i+1]-'0')
				i++
			}
			q.PushOperand(val)
		} else if c == ' ' {
			continue
		} else if c == '(' {
			q.PushOperator(c)
		} else if c == ')' {
			for q.PeekOperator() != '(' {
				q.PushOperand(q.operate())
			}
			q.PopOperator()
		} else {
			for !q.isOperatorEmpty() && comparePrecedence(c, q.PeekOperator()) <= 0 {
				q.PushOperand(q.operate())
			}
			q.PushOperator(c)
		}
	}

	for !q.isOperatorEmpty() {
		q.PushOperand(q.operate())
	}

	return q.PopOperand()
}

func comparePrecedence(a, b byte) int {
	return precedenceMap[a] - precedenceMap[b]
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
