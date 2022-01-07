package Medium

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

func (this *Solution) GetRandom() int {
	cnt, node, candidate := 0, this.head, this.head

	for node != nil {
		cnt++
		if rand.Intn(cnt) == 0 {
			candidate = node
		}
		node = node.Next
	}

	return candidate.Val
}
