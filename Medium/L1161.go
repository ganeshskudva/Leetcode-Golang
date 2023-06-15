package Medium

func maxLevelSum(root *TreeNode) int {
	var q []*TreeNode
	maximum, maxLevel, lvl := math.MinInt32, -1, 1

	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		sum := 0
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			sum += node.Val

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum > maximum {
			maximum = sum
			maxLevel = lvl
		}
		lvl++
	}

	return maxLevel
}
