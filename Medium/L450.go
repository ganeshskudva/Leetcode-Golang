package Medium

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, root.Val)
	}

	return root
}

func findMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}

	return root
}
