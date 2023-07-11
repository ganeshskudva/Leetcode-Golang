package Medium

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var (
		res    []int
		find   func(root, target *TreeNode)
		search func(root *TreeNode, distance int)
	)
	mp := make(map[*TreeNode]int)

	find = func(root, target *TreeNode) {
		if root == nil {
			return
		}

		if root == target {
			mp[root] = 0
			return
		}

		find(root.Left, target)
		if val, ok := mp[root.Left]; ok {
			mp[root] = val + 1
			return
		}

		find(root.Right, target)
		if val, ok := mp[root.Right]; ok {
			mp[root] = val + 1
			return
		}
	}

	search = func(root *TreeNode, distance int) {
		if root == nil {
			return
		}

		if val, ok := mp[root]; ok {
			distance = val
		}

		if distance == k {
			res = append(res, root.Val)
		}

		search(root.Left, distance+1)
		search(root.Right, distance+1)
	}

	find(root, target)
	search(root, 0)

	return res
}
