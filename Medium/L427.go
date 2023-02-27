package Medium

func construct(grid [][]int) *Node {
	var solve func(x1, x2, y1, y2 int) *Node
	solve = func(x1, x2, y1, y2 int) *Node {
		if x1 == x2 {
			val := false
			if grid[x1][y1] == 1 {
				val = true
			}
			return &Node{
				Val:    val,
				IsLeaf: true,
			}
		}

		rowMid, colMid := (x1+x2)/2, (y1+y2)/2
		topLeft := solve(x1, rowMid, y1, colMid)
		topRight := solve(x1, rowMid, colMid+1, y2)
		bottomLeft := solve(rowMid+1, x2, y1, colMid)
		bottomRight := solve(rowMid+1, x2, colMid+1, y2)

		if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf &&
			topRight.Val == topLeft.Val && topRight.Val == bottomLeft.Val && topRight.Val == bottomRight.Val {
			return &Node{
				Val:    topLeft.Val,
				IsLeaf: true,
			}
		} else {
			return &Node{
				TopLeft:     topLeft,
				TopRight:    topRight,
				BottomLeft:  bottomLeft,
				BottomRight: bottomRight,
			}
		}
	}

	return solve(0, len(grid)-1, 0, len(grid[0])-1)
}
