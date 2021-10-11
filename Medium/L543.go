package Medium

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	maxDepth(root, &max)
	return max
}

func maxDepth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left, right := maxDepth(root.Left, max), maxDepth(root.Right, max)

	if left+right > *max {
		*max = left + right
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
