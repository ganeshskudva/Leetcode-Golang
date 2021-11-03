package Medium

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, currSum int) int {
	if root == nil {
		return 0
	}

	currSum = currSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return currSum
	}

	leftSum, rightSum := dfs(root.Left, currSum), dfs(root.Right, currSum)

	return leftSum + rightSum
}
