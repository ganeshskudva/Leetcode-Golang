/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package Easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
		return 0
	}

	var q []*TreeNode
	q = append(q, root)
	sum := 0

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if node.Left != nil && isLeaf(node.Left) {
				sum += node.Left.Val
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return sum
}

func isLeaf(root *TreeNode) bool {
	return root.Right == nil && root.Left == nil
}
