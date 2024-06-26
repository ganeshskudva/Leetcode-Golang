package Medium

func balanceBST(root *TreeNode) *TreeNode {
	var (
		sortedLst []*TreeNode
		inorder func(node *TreeNode) 
		createTree func(start, end int) *TreeNode
	)
	
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		sortedLst = append(sortedLst, node)
		inorder(node.Right)
	}
	
	createTree = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := start + (end-start)/2
		newRoot := sortedLst[mid]
		newRoot.Left = createTree(start, mid - 1)
		newRoot.Right = createTree(mid + 1, end)
		return newRoot
	}
	
	inorder(root)
	return createTree(0, len(sortedLst) - 1)
}
