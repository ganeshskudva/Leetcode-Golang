package Medium

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
	}

	preMiddle, preCurrent := p1, p1.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}

	p1, p2 = head, preMiddle.Next
	for p1 != preMiddle {
		preMiddle.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = preMiddle.Next
	}
}
