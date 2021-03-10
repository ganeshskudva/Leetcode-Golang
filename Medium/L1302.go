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
	q := make([]*TreeNode, 1, 10)
	q[0] = root
	
	for len(q) > 0 {
		size := len(q)
		sum = 0
		for ; size > 0 ; size-- {
			node := q[0]
			q = q[1:]
			
			if node.Left == nil && node.Right == nil {
				sum += node.Val
			} else {
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
		}
	}
	
	return sum
}
