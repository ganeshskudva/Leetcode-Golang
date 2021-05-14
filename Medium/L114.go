package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    	for root != nil {
		if root.Left != nil && root.Right != nil{
			t := root.Left
			for t.Right != nil {
				t = t.Right
			}
			t.Right = root.Right
		}
		
		if root.Left != nil {
			root.Right = root.Left
		}
		root.Left = nil
		root = root.Right
	}
}
