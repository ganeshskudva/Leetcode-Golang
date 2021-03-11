package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var res []int
	if root1 == nil && root2 == nil {
		return res
	}

	var st1, st2 []*TreeNode
	pushLeft(&st1, root1)
	pushLeft(&st2, root2)

	for len(st1) != 0 || len(st2) != 0 {
		var st *[]*TreeNode
		if len(st1) == 0 {
			st = &st2
		} else if len(st2) == 0 {
			st = &st1
		} else {
			if st1[len(st1)-1].Val < st2[len(st2)-1].Val {
				st = &st1
			} else {
				st = &st2
			}
		}

		curr := (*st)[len(*st)-1]
		*st = (*st)[:len(*st)-1]
		res = append(res, curr.Val)
		pushLeft(st, curr.Right)
	}

	return res
}

func pushLeft(st *[]*TreeNode, node *TreeNode) {
	for node != nil {
		*st = append(*st, node)
		node = node.Left
	}
}
