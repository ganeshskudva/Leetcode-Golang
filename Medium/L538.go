package Medium

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convert(root, &sum)
	return root
}

func convert(cur *TreeNode, sum *int) {
	if cur == nil {
		return
	}

	convert(cur.Right, sum)
	cur.Val += *sum
	*sum = cur.Val
	convert(cur.Left, sum)
}
