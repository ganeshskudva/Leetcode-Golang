package Medium 

type MyCalendar struct {
	root *treeNode
}

type treeNode struct {
	start, end  int
	left, right *treeNode
}

func Constructor() MyCalendar {
	return MyCalendar{
		root: nil,
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if end < start || start < 0 {
		return false
	}
 
    return insert(&this.root, start, end)
}

func insert(root **treeNode, start, end int) bool {
	if *root == nil {
		*root = &treeNode{
			start: start,
			end:   end,
			left:  nil,
			right: nil,
		}
		return true
	}

	if (*root).start >= end {
		return insert(&(*root).left, start, end)
	} else if (*root).end <= start {
		return insert(&(*root).right, start, end)
	}
	return false
}
