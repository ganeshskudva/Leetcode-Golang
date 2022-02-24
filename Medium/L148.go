package Medium

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// step 1. cut the list to two halves
	var prev *ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		prev, slow, fast = slow, slow.Next, fast.Next.Next
	}

	prev.Next = nil

	// step 2. sort each half
	l1, l2 := sortList(head), sortList(slow)

	// step 3. merge l1 and l2
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	l := &ListNode{}
	p := l

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next, l1 = l1, l1.Next
		} else {
			p.Next, l2 = l2, l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}

	return l.Next
}
