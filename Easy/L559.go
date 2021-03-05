package Easy

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    	if root == nil {
		return 0
	}

	max := 0
	q := make([]*Node, 1, 10)
	q[0] = root

	for len(q) > 0 {
		max++
		size := len(q)
		for i:= 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			for _, v := range n.Children {
				q = append(q, v)
			}
		}
	}

	return max
}
