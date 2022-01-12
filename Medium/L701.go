package Medium

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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return NewTreeNode(val)
	}
	
	helper(root, val)
	return root
}

func helper(root *TreeNode, val int) {
	if root.Val < val && root.Right == nil {
		root.Right = NewTreeNode(val)
	} else if root.Val > val && root.Left == nil {
		root.Left = NewTreeNode(val)
	} else if root.Val > val {
		helper(root.Left, val)
	} else {
		helper(root.Right, val)
	}
}
