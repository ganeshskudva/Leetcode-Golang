package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	dfs(root, root.Val, root.Val, &res)
	
	return res
}

func dfs(root *TreeNode, minimum, maximum int, res *int) {
	if root == nil {
		return
	}
	minimum, maximum = min(root.Val, minimum), max(root.Val, maximum)
	*res = max(*res, maximum - minimum)
	
	dfs(root.Left, minimum, maximum, res)
	dfs(root.Right, minimum, maximum, res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	
	return b
}
