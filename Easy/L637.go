package Easy

func averageOfLevels(root *TreeNode) []float64 {
	var (
		res []float64
		q   []*TreeNode
	)
	if root == nil {
		return res
	}

	q = append(q, root)
	for len(q) != 0 {
		n := len(q)
		sum := 0
		for i := 0; i < n; i++ {
			cur := q[0]
			sum += cur.Val
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append(res, float64(sum)/float64(n))
	}
	return res
}
