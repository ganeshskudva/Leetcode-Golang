package Easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var (
		root1Leaves, root2Leaves string
		solve                    func(root *TreeNode, rootLeaves *string)
	)

	solve = func(root *TreeNode, rootLeaves *string) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			*rootLeaves = fmt.Sprintf("%s%d#", *rootLeaves, root.Val)
		}
		solve(root.Left, rootLeaves)
		solve(root.Right, rootLeaves)
	}

	solve(root1, &root1Leaves)
	solve(root2, &root2Leaves)

	return root1Leaves == root2Leaves
}
