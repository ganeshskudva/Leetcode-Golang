package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    	if root == nil {
		return 0
	}
	
	sum := 0
	var q []*TreeNode
	q = append(q, root)
	
	for len(q) > 0 {
		size := len(q)
		sum = 0
		for ; size > 0; size-- {
			n := q[0]
			q = q[1:]
			
			if n.Left == nil && n.Right == nil {
				sum += n.Val
			}
			
			if n.Left != nil {
				q = append(q, n.Left)
			}
			
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
	}
	
	return sum
}
