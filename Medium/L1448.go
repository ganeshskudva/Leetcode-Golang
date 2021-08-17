package Medium

func goodNodes(root *TreeNode) int {
	cnt := 0
	if root == nil {
		return cnt
	}
	dfs(root, root.Val, &cnt)
	return cnt
}

func dfs(root *TreeNode, val int, cnt *int) {
	if root == nil {
		return
	}

	if root.Val >= val {
		*cnt = (*cnt) + 1
	}

	max := val
	if val < root.Val {
		max = root.Val
	}

	dfs(root.Left, max, cnt)
	dfs(root.Right, max, cnt)
}
