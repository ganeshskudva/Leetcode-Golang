package Medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
        kthFromStart, kthFromEnd := head, head
    for node := head; node != nil; node = node.Next {
        k--
        if k == 0 {
            kthFromStart = node
        }
        if k < 0 {
            kthFromEnd = kthFromEnd.Next
        }
    }
    kthFromStart.Val, kthFromEnd.Val = kthFromEnd.Val, kthFromStart.Val
    return head
}
