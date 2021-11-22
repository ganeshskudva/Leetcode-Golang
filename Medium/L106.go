package Medium

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	mp := make(map[int]int)
	for i := range inorder {
		mp[inorder[i]] = i
	}

	return helper(postorder, 0, len(postorder)-1, inorder, 0, len(inorder)-1, mp)
}

func helper(postOrder []int, postStart, postEnd int, inOrder []int, inStart, inEnd int, mp map[int]int) *TreeNode {
	if postStart > postEnd || inStart > inEnd {
		return nil
	}

	root := &TreeNode{
		Val: postOrder[postEnd],
	}
	inRoot := mp[root.Val]
	numsLeft := inRoot - inStart

	root.Left = helper(postOrder, postStart, postStart+numsLeft-1, inOrder, inStart, inRoot-1, mp)
	root.Right = helper(postOrder, postStart+numsLeft, postEnd-1, inOrder, inRoot+1, inEnd, mp)

	return root
}
