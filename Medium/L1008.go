package Medium

func bstFromPreorder(preorder []int) *TreeNode {
	idx := 0
	return dfs(preorder, math.MinInt64, math.MaxInt64, &idx)
}

func dfs(preorder []int, low, high int, idx *int) *TreeNode {
	if *idx >= len(preorder) {
		return nil
	}
	if preorder[*idx] < low || preorder[*idx] > high {
		return nil
	}

	root := &TreeNode{
		Val: preorder[*idx],
	}
	*idx += 1
	root.Left, root.Right = dfs(preorder, low, root.Val, idx), dfs(preorder, root.Val, high, idx)

	return root
}
