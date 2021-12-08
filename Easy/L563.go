package Easy

func findTilt(root *TreeNode) int {
	res := 0

	postOrder(root, &res)
	return res
}

func postOrder(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left, right := postOrder(root.Left, res), postOrder(root.Right, res)

	*res += abs(left - right)

	return left + right + root.Val
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
