package Easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
	var (
		sb strings.Builder
		solve func(root *TreeNode)
	)
	if root == nil {
		return ""
	}
	
	solve = func(root *TreeNode) {
		sb.WriteString(fmt.Sprintf("%d", root.Val))
		if root.Left == nil && root.Right == nil {
			return
		}
		
		if root.Left != nil {
			sb.WriteString("(")
			solve(root.Left)
			sb.WriteString(")")
		}
		
		if root.Right != nil {
			if root.Left == nil {
				sb.WriteString("()")
			}
			
			sb.WriteString("(")
			solve(root.Right)
			sb.WriteString(")")
		}
	}
	
	solve(root)
	return sb.String()
}
