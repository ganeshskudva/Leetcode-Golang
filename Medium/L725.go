package Medium

func splitListToParts(head *ListNode, k int) []*ListNode {
	n, curr := 0, head
	for curr != nil {
		n += 1
		curr = curr.Next
	}
	m, r, j := n/k, n%k, 0
	if m == 0 {
		r -= k
	}

	res := make([]*ListNode, k)
	curr = head
	for curr != nil {
		res[j] = curr
		j++
		for i := 1; i < m; i++ {
			curr = curr.Next
		}
		if r > 0 {
			curr = curr.Next
			r -= 1
		}
		tmp := curr.Next
		curr.Next = nil
		curr = tmp
	}

	return res
}
