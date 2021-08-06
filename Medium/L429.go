package Medium

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var q []*Node
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		var tmp []int
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			tmp = append(tmp, n.Val)

			for _, v := range n.Children {
				q = append(q, v)
			}
		}
		res = append(res, tmp)
	}

	return res
}
