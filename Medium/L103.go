package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q, root)
	normalOrder := false

	for len(q) > 0 {
		size := len(q)
		var lvlVals []int
		normalOrder = !normalOrder
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]

			if normalOrder {
				lvlVals = append(lvlVals, cur.Val)
			} else {
				lvlVals = append([]int{cur.Val}, lvlVals...)
			}
			
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append(res, lvlVals)
	}
	
	return res
}
