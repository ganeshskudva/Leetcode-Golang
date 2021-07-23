package Medium

func pruneTree(root *TreeNode) *TreeNode {
	return dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left = dfs(root.Left)
	root.Right = dfs(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
