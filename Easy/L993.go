package Easy

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	q := make([]*TreeNode, 1, 10)
	q[0] = root
	var xExist, yExist bool

	for len(q) > 0 {
		size := len(q)
		xExist, yExist = false, false
		for ; size > 0; size-- {
			n := q[0]
			q = q[1:]

			if n.Val == x {
				xExist = true
			}
			if n.Val == y {
				yExist = true
			}

			if n.Left != nil && n.Right != nil {
				if n.Left.Val == x && n.Right.Val == y {
					return false
				}

				if n.Left.Val == y && n.Right.Val == x {
					return false
				}
			}

			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		if xExist && yExist {
			return true
		}
	}

	return false
}
