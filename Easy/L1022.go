package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var tmp string
	sum := 0

	dfs(root, &sum, &tmp)

	return sum
}

func dfs(root *TreeNode, sum *int, tmp *string) {
	if root == nil {
		return
	}
	*tmp += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		val, _ := strconv.ParseInt(*tmp, 2, 64)
		*sum += int(val)
	} else {
		dfs(root.Left, sum, tmp)
		dfs(root.Right, sum, tmp)
	}

	*tmp = (*tmp)[:len(*tmp)-1]
}
