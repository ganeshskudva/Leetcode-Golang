package Easy

func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	postOrder(root, &res)

	return res
}

func postOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}

	postOrder(root.Left, list)
	postOrder(root.Right, list)
	*list = append(*list, root.Val)
}
