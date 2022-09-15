func pseudoPalindromicPaths(root *TreeNode) int {
	res, mp := 0, make([]int, 10)
	findPesudoPalindromUtil(&res, &mp, root)

	return res
}

func findPesudoPalindromUtil(res *int, mp *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	(*mp)[root.Val] = (*mp)[root.Val] + 1
	if root.Left == nil && root.Right == nil {
		if isPalindrome(mp) {
			*res++
		}
	}

	findPesudoPalindromUtil(res, mp, root.Left)
	findPesudoPalindromUtil(res, mp, root.Right)

	(*mp)[root.Val] = (*mp)[root.Val] - 1
}

func isPalindrome(mp *[]int) bool {
	miss := 0

	for i := 0; i <= 9; i++ {
		if (*mp)[i]%2 != 0 {
			miss++
		}
		if miss > 1 {
			return false
		}
	}

	return true
}
