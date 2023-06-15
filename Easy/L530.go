package Easy

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getMinimumDifference(root *TreeNode) int {
	var (
		solve func(root *TreeNode) int
	)
	minimum, prev := math.MaxInt32, -1

	solve = func(root *TreeNode) int {
		if root == nil {
			return minimum
		}

		solve(root.Left)

		if prev != -1 {
			minimum = min(minimum, root.Val-prev)
		}
		prev = root.Val

		solve(root.Right)

		return minimum
	}

	return solve(root)
}
