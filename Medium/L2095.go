package Medium

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}
