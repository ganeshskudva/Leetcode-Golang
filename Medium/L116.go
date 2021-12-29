// DFS solution
func connect(root *Node) *Node {
	dfs(root, nil)
	return root
}

func dfs(curr, next *Node) {
	if curr == nil {
		return
	}

	curr.Next = next
	dfs(curr.Left, curr.Right)

	if curr.Next != nil {
		dfs(curr.Right, curr.Next.Left)
	} else {
		dfs(curr.Right, nil)
	}
}

// BFS solution
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	var q []*Node
	q = append(q, root)

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			curr := q[0]
			q = q[1:]

			if i != size-1 {
				curr.Next = q[0]
			}

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
	}

	return root
}
