package Easy

func maxDepth(root *TreeNode) int {
	maxLevel := 0
	dfs(root, 1, &maxLevel)

	return maxLevel
}

func dfs(root *TreeNode, currLevel int, maxLevel *int) {
	if root == nil {
		return
	}

	if currLevel > *maxLevel {
		*maxLevel = currLevel
	}

	dfs(root.Left, currLevel+1, maxLevel)
	dfs(root.Right, currLevel+1, maxLevel)
}
