package Medium

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	mp := make(map[int]int)
	for i := range inorder {
		mp[inorder[i]] = i
	}

	return solve(preorder, inorder, mp, 0, len(preorder)-1, 0, len(inorder)-1)
}

func solve(preorder, inorder []int, mp map[int]int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	root := &TreeNode{
		Val:   preorder[preStart],
		Left:  nil,
		Right: nil,
	}
	inRoot := mp[root.Val]
	numLeft := inRoot - inStart

	root.Left = solve(preorder, inorder, mp, preStart+1, preStart+numLeft, inStart, inRoot-1)
	root.Right = solve(preorder, inorder, mp, preStart+numLeft+1, preEnd, inRoot+1, inEnd)

	return root
}
