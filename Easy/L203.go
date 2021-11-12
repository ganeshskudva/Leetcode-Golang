package Easy

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElements(head.Next, val)

	if head.Val == val {
		return head.Next
	}

	return head
}
