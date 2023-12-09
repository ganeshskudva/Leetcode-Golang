package Easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var (
		res   []int
		solve func(root *TreeNode)
	)

	solve = func(root *TreeNode) {
		if root == nil {
			return
		}

		solve(root.Left)
		res = append(res, root.Val)
		solve(root.Right)
	}

	solve(root)
	return res
}
