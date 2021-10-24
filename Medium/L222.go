package Medium

func countNodes(root *TreeNode) int {
	cnt := 0
	if root == nil {
		return cnt
	}

	var q []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)

		for idx := 0; idx < size; idx++ {
			node := q[0]
			q = q[1:]
			cnt++

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return cnt
}
