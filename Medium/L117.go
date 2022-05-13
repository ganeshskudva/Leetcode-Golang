package Medium

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	var q []*Node
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		tmp := make([]*Node, len(q))
		copy(tmp, q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if i != size-1 {
				node.Next = tmp[i+1]
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return root
}
