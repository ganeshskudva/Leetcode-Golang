package Easy

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, targetSum)
}

func dfs(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	} else if root.Right == nil && root.Left == nil && sum == root.Val {
		return true
	} else {
		sum -= root.Val
		return dfs(root.Left, sum) || dfs(root.Right, sum)
	}
}
