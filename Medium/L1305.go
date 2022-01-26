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
	var q1, q2, ans []int

	inorder(root1, &q1)
	inorder(root2, &q2)

	for len(q1) > 0 || len(q2) > 0 {
		if len(q2) == 0 {
			ans = append(ans, q1[0])
			q1 = q1[1:]
		} else if len(q1) == 0 {
			ans = append(ans, q2[0])
			q2 = q2[1:]
		} else {
			if q1[0] < q2[0] {
				ans = append(ans, q1[0])
				q1 = q1[1:]
			} else {
				ans = append(ans, q2[0])
				q2 = q2[1:]
			}
		}
	}

	return ans
}

func inorder(root *TreeNode, q *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, q)
	*q = append(*q, root.Val)
	inorder(root.Right, q)
}
