package Medium

type Node struct {
	ch  byte
	tot int
}

func NewNode(ch byte) Node {
	return Node{
		ch:  ch,
		tot: 1,
	}
}
func removeDuplicates(s string, k int) string {
	var st []Node

	for i, _ := range s {
		if len(st) > 0 && st[len(st)-1].ch == s[i] {
			st[len(st)-1].tot++
		} else {
			st = append(st, NewNode(s[i]))
		}

		if st[len(st)-1].tot == k {
			st = st[:len(st)-1]
		}
	}

	var sb strings.Builder
	for i, _ := range st {
		for j := 0; j < st[i].tot; j++ {
			sb.WriteByte(st[i].ch)
		}
	}

	return sb.String()
}
