package Easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	
	var solve func(left, right *TreeNode) bool
	solve = func(left, right *TreeNode) bool {
		if left == nil || right == nil {
			return left == right
		}
		if left.Val != right.Val {
			return false
		}
		
		return solve(left.Left, right.Right) && solve(left.Right, right.Left)
	}
	
	return solve(root.Left, root.Right)
}
