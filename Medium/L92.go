package Medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || (left == right) {
		return head
	}

	var start, end, startPrev, endNext *ListNode
	curr := head
	for i := 1; i < left; i++ {
		startPrev = curr
		curr = curr.Next
	}

	start = curr
	curr = head
	for i := 1; i < right; i++ {
		curr = curr.Next
	}
	end = curr
	endNext = curr.Next
	reverse(start, end, left, right)
	if startPrev == nil {
		if endNext == nil {
			return end
		}

		start.Next = endNext
		return end
	}

	startPrev.Next = end
	start.Next = endNext

	return head
}

func reverse(start, end *ListNode, m, n int) {
	var prev, tmp *ListNode
    curr := start

	for i := m; i <= n; i++ {
		tmp = curr
		curr = curr.Next
		tmp.Next = prev
		prev = tmp
	}
}
