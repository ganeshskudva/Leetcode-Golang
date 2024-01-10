package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
	var (
		amount int
		solve  func(node *TreeNode, res *int) int
	)

	solve = func(node *TreeNode, res *int) int {
		if node == nil {
			return 0
		}

		left, right := solve(node.Left, res), solve(node.Right, res)

		if node.Val == start {
			*res = max(left, right)
			return -1
		} else if left >= 0 && right >= 0 {
			return max(left, right) + 1
		}

		*res = max(*res, abs(left-right))
		return min(left, right) - 1
	}

	solve(root, &amount)
	return amount
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
