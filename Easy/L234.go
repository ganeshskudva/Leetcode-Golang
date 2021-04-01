package Easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var tmp *ListNode

func isPalindrome(head *ListNode) bool {
	tmp = head

	return check(head)
}

func check(p *ListNode) bool {
	if p == nil {
		return true
	}

	isPal := false
	if check(p.Next) && tmp.Val == p.Val {
		isPal = true
	}
	tmp = tmp.Next

	return isPal
}
