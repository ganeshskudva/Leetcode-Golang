package Medium

func smallestFromLeaf(root *TreeNode) string {
	var (
		tmp    string
		isLeaf func(node *TreeNode) bool
		solve  func(node *TreeNode, s string) string
	)

	isLeaf = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		return node.Left == nil && node.Right == nil
	}

	solve = func(node *TreeNode, res string) string {
		if node == nil {
			return "|"
		}

		res = fmt.Sprintf("%c", byte('a'+node.Val)) + res
		if isLeaf(node) {
			return res
		}

		left, right := solve(node.Left, res), solve(node.Right, res)
		if strings.Compare(left, right) < 0 {
			return left
		}

		return right
	}

	return solve(root, tmp)
}
