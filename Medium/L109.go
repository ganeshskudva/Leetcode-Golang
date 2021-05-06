package Medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	
	if head.Next == nil {
		return NewTreeNode(head.Val)
	}
	
	var prev *ListNode
	slow, fast := head, head
	
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	prev.Next = nil
	
	root := NewTreeNode(slow.Val)
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)
	
	return root
}
