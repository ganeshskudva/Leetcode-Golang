package Medium

func pathSum(root *TreeNode, targetSum int) int {
	mp := make(map[int]int)
	mp[0] = 1

	return solve(root, 0, targetSum, mp)
}

func solve(root *TreeNode, currSum, target int, mp map[int]int) int {
	if root == nil {
		return 0
	}

	currSum += root.Val
	res := 0
	if val, ok := mp[currSum-target]; ok {
		res = val
	}
	if val, ok := mp[currSum]; ok {
		mp[currSum] = val + 1
	} else {
		mp[currSum] = 1
	}

	res += solve(root.Left, currSum, target, mp) + solve(root.Right, currSum, target, mp)
	if val, ok := mp[currSum]; ok {
		mp[currSum] = val - 1
	}

	return res
}
