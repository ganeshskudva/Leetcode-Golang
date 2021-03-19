package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var tmp []int
	if root == nil {
		return res
	}

	bt(root, targetSum, &res, &tmp)

	return res
}

func bt(root *TreeNode, tgt int, res *[][]int, tmp *[]int) {
	if root == nil {
		return
	}

	*tmp = append(*tmp, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == tgt {
		cp := make([]int, len(*tmp))
		copy(cp, *tmp)
		*res = append(*res, cp)
	} else {
		bt(root.Left, tgt-root.Val, res, tmp)
		bt(root.Right, tgt-root.Val, res, tmp)
	}

	*tmp = (*tmp)[:len(*tmp)-1]
}
