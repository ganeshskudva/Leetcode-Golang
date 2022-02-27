package Medium

func widthOfBinaryTree(root *TreeNode) int {
	maxWidth := 0
	var lst []int
	dfs(root, 1, 0, &lst, &maxWidth)

	return maxWidth
}

func dfs(root *TreeNode, id, depth int, lst *[]int, maxWidth *int) {
	if root == nil {
		return
	}
	if depth >= len(*lst) {
		*lst = append(*lst, id)
	}
	if (id + 1 - (*lst)[depth]) > *maxWidth {
		*maxWidth = id + 1 - (*lst)[depth]
	}

	if root.Left != nil {
		dfs(root.Left, id*2, depth+1, lst, maxWidth)
	}
	if root.Right != nil {
		dfs(root.Right, id*2+1, depth+1, lst, maxWidth)
	}
}
