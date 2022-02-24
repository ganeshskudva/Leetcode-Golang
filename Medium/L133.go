package Medium

func cloneGraph(node *Node) *Node {
	return dfs(node, make(map[int]*Node))
}

func dfs(node *Node, mp map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if val, ok := mp[node.Val]; ok {
		return val
	}

	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	mp[newNode.Val] = newNode
	for _, n := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, dfs(n, mp))
	}

	return newNode
}
