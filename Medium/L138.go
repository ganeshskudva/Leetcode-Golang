package Medium

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	mp, node := make(map[*Node]*Node), head
	// loop 1. copy all the nodes
	for node != nil {
		mp[node] = &Node{Val: node.Val}
		node = node.Next
	}

	node = head
	// loop 2. assign next and random pointers
	for node != nil {
		mp[node].Next = mp[node.Next]
		mp[node].Random = mp[node.Random]
		node = node.Next
	}

	return mp[head]
}
