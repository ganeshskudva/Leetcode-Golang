package Easy

func invertTree(root *TreeNode) *TreeNode {
	return dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left, right := root.Left, root.Right
	root.Left = dfs(right)
	root.Right = dfs(left)

	return root
}
