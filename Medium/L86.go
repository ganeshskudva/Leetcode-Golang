package Medium

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func partition(head *ListNode, x int) *ListNode {
	dummyHead, greatHead := NewListNode(-1), NewListNode(-1)

	small, great, point := dummyHead, greatHead, head
	dummyHead.Next = head

	for point != nil {
		if point.Val < x {
			small.Next = point
			small = small.Next
			point = point.Next
		} else {
			great.Next = point
			great = great.Next
			point = point.Next
		}
	}
	great.Next = nil
	small.Next = greatHead.Next

	return dummyHead.Next
}
