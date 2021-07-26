package Easy

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	head := solve(nums, 0, len(nums) - 1)
	return head
}

func solve(nums []int, lo, hi int) *TreeNode{
	if lo > hi {
		return nil
	}
	
	mid := lo + (hi - lo)/2
	node := new(TreeNode)
	node.Val = nums[mid]
	node.Left, node.Right = solve(nums, lo, mid -1), solve(nums, mid + 1, hi)
	
	return node
}
