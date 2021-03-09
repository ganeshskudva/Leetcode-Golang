package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
    if root == nil {
		return root
	}
	
	if d == 1 {
		newRoot := TreeNode{
			Val:   v,
			Left:  nil,
			Right: nil,
		}
		
		newRoot.Left = root
		
		return &newRoot
	}

	var node *TreeNode
	lvl := 0
	q := make([]*TreeNode, 1, 10)
	q[0] = root

	for lvl < d {
		size := len(q)
		lvl++
		for ; size > 0; size-- {
			node = q[0]
			q = q[1:]

			if lvl == d-1 {
				Left := TreeNode{
					Val:   v,
					Left:  nil,
					Right: nil,
				}

				Right := TreeNode{
					Val:   v,
					Left:  nil,
					Right: nil,
				}

				Left.Left = node.Left
				Right.Right = node.Right

				node.Left = &Left
				node.Right = &Right

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

	return root
}
