package Easy

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return dfs(root1, root2)
}

func dfs(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	var root *TreeNode

	if t1 != nil && t2 == nil {
		root = &TreeNode{Val: t1.Val}
	} else if t1 == nil && t2 != nil {
		root = &TreeNode{Val: t2.Val}
	} else {
		root = &TreeNode{Val: t1.Val + t2.Val}
	}

	root.Left = dfs(t1.Left, t2.Left)
	root.Right = dfs(t1.Right, t2.Right)

	return root
}
