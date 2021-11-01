package Medium

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	p := root
	for p != nil {
		if p.Child == nil {
			p = p.Next
			continue
		}
		tmp := p.Child
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = p.Next
		if p.Next != nil {
			p.Next.Prev = tmp
		}
		p.Next = p.Child
		p.Child.Prev = p
		p.Child = nil
	}

	return root
}
