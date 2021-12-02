package Medium

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd, even, evenHead := head, head.Next, head.Next

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead

	return head
}
