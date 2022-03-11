package Medium

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	listNum, tail := 1, head

	for tail.Next != nil {
		listNum++
		tail = tail.Next
	}
	tail.Next = head
	newHeadIndex := listNum - k%listNum

	for i := 0; i < newHeadIndex; i++ {
		tail = tail.Next
	}

	head = tail.Next
	tail.Next = nil

	return head
}
