package Easy

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, root, k)
}

func dfs(root, cur *TreeNode, k int) bool {
	if cur == nil {
		return false
	}

	return search(root, cur, k-cur.Val) || dfs(root, cur.Left, k) || dfs(root, cur.Right, k)
}

func search(root, curr *TreeNode, val int) bool {
	if root == nil {
		return false
	}

	return (root.Val == val) && (root != curr) || (root.Val < val) && search(root.Right, curr, val) || (root.Val > val) && search(root.Left, curr, val)
}
