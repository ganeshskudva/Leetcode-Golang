package Easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
	var solve func(root *TreeNode) bool
	mp := map[int]bool{
		0: false,
		1: true,
	}

	solve = func(node *TreeNode) bool {
		if node.Right == nil && node.Left == nil {
			return mp[node.Val]
		}
		
		var left, right bool
		if node.Left != nil {
			left = solve(node.Left)
		}
		if node.Right != nil {
			right = solve(node.Right)
		}
		
		if node.Val == 2 {
			return left || right
		}
		
		return left && right
	}
	
	return solve(root)
}
