package Hard

func maxPathSum(root *TreeNode) int {
	maxi := math.MinInt32

	solve(root, &maxi)

	return maxi
}

func solve(root *TreeNode, maximum *int) int {
	if root == nil {
		return 0
	}

	left, right := max(solve(root.Left, maximum), 0), max(solve(root.Right, maximum), 0)

	*maximum = max(*maximum, root.Val+right+left)

	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
