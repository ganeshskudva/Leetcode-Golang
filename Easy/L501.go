package Easy

func findMode(root *TreeNode) []int {
	var (
		modes []int
		solve func(node *TreeNode, count *int)
	)
	prev, cnt, maxCnt := root, 0, math.MinInt32

	solve = func(node *TreeNode, count *int) {
		if node == nil {
			return
		}
		solve(node.Left, count)
		if prev.Val == node.Val {
			*count++
		} else {
			*count = 1
		}

		if *count == maxCnt {
			modes = append(modes, node.Val)
		} else if *count > maxCnt {
			modes = nil
			modes = append(modes, node.Val)
			maxCnt = *count
		}

		prev = node
		solve(node.Right, count)
	}

	solve(root, &cnt)
	return modes
}
