package Easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Iterative Solution
func preorderTraversal(root *TreeNode) []int {
	var (
		res []int
		st []*TreeNode
	)
	st = append(st, root)
	
	for len(st) > 0 {
		node := st[len(st) - 1]
		st = st[:len(st) - 1]
		if node != nil {
			res = append(res, node.Val)
			st = append(st, node.Right)
			st = append(st, node.Left)
		}
	}
	
	return res
}


// Recursive Solution
func preorderTraversal(root *TreeNode) []int {
	var res []int

	solve(root, &res)
	return res
}

func solve(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	solve(root.Left, res)
	solve(root.Right, res)
}
