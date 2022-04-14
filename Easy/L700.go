package Easy

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || val == root.Val {
		return root
	}
	
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	
	return searchBST(root.Right, val)
}
