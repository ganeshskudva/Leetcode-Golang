package Medium

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	var st []*TreeNode
	st = append(st, root)

	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		for node != nil {
			st = append(st, node)
			node = node.Left
		}

		if len(st) > 0 {
			node = st[len(st)-1]
			st = st[:len(st)-1]

			k--
			if k == 0 {
				return node.Val
			}
			st = append(st, node.Right)
		}
	}

	return 0
}
