package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
    var res []int
	if root == nil {
		return res
	}

	i := 0
	st := make([]*TreeNode, len(voyage))
	st[0] = root

	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]

		if node == nil {
			continue
		}

		if node.Val != voyage[i] {
			return []int{-1}
		}
		i++
		if node.Right != nil && node.Right.Val == voyage[i] {
			if node.Left != nil {
				res = append(res, node.Val)
			}

			st = append(st, node.Left)
			st = append(st, node.Right)
		} else {
			st = append(st, node.Right)
			st = append(st, node.Left)
		}
	}

	return res
}
