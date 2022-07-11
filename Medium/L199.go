package Medium

func rightSideView(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)

		for size > 0 {
			size--
			cur := q[0]
			q = q[1:]

			if size == 0 {
				res = append(res, cur.Val)
			}

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}

	return res
}
