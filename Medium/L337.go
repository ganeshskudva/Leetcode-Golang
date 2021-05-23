package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	return solve(root, make(map[*TreeNode]int))
}

func solve(root *TreeNode, dp map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if v, ok := dp[root]; ok {
		return v
	}

	val := 0
	if root.Left != nil {
		val += solve(root.Left.Left, dp) + solve(root.Left.Right, dp)
	}
	if root.Right != nil {
		val += solve(root.Right.Left, dp) + solve(root.Right.Right, dp)
	}

	val = max(val+root.Val, solve(root.Left, dp)+solve(root.Right, dp))
	dp[root] = val

	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
