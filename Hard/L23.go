package Hard

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var (
		merge func(l1, l2 *ListNode) *ListNode
		partition func(start, end int) *ListNode
	)
	
	merge = func(l1, l2 *ListNode) *ListNode {
		if l1 == nil {
			return l2
		}
		
		if l2 == nil {
			return l1
		}
		
		if l1.Val < l2.Val {
			l1.Next = merge(l1.Next, l2)
			return l1
		}
		
		l2.Next = merge(l1, l2.Next)
		return l2
	}
	
	partition = func(start, end int) *ListNode {
		if start == end {
			return lists[start]
		}
		
		if start >= end {
			return nil
		}
		
		q := (start + end) / 2
		l1, l2 := partition(start, q), partition(q + 1, end)
		
		return merge(l1, l2)
	}
	
	return partition(0, len(lists) - 1)
}
