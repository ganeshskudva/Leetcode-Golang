package Medium

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minDiffInBST(root *TreeNode) int {
	var prev *TreeNode
	minimum := math.MaxInt32

	solve := func(root *TreeNode, prev **TreeNode) {}
	solve = func(root *TreeNode, prev **TreeNode) {
		if root == nil {
			return
		}

		solve(root.Left, prev)
		if prev != nil && *prev != nil {
			minimum = min(minimum, root.Val-(*prev).Val)
		}
		*prev = root
		solve(root.Right, prev)
	}

	solve(root, &prev)
	return minimum
}
