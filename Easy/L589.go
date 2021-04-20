package Easy

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	var res []int 
	dfs(root, &res)
	return res
}

func dfs(root *Node, res *[]int) {
	if root == nil {
		return
	}
	
	*res = append(*res, root.Val)
	for _, v := range root.Children {
		dfs(v, res)
	}
}
