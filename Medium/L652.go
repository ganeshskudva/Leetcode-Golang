package Medium

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	curID := 1
	serialToID, idToCount := make(map[string]int), make(map[int]int)
	var (
		res   []*TreeNode
		solve func(root *TreeNode) int
	)
	solve = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftID, rightID := solve(root.Left), solve(root.Right)
		curSerial := fmt.Sprintf("%d-%d-%d", leftID, root.Val, rightID)
		serialID := curID
		if v, ok := serialToID[curSerial]; ok {
			serialID = v
		}
		if serialID == curID {
			curID++
		}
		serialToID[curSerial] = serialID
		idToCount[serialID]++

		if idToCount[serialID] == 2 {
			res = append(res, root)
		}
		return serialID
	}

	solve(root)
	return res
}
