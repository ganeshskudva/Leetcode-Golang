package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}
	
	dfs(root, nil, nil, &sum)
	
	return sum
}

func dfs(root, parent, grand *TreeNode, sum *int)  {
	if root == nil {
		return
	}
	
	if grand != nil && grand.Val%2 == 0 {
		*sum += root.Val
	}
	
	dfs(root.Left, root, parent, sum)
	dfs(root.Right, root, parent, sum)
	
}
