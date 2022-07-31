package Medium

type NumArray struct {
	root *SegmentTree
}

type SegmentTree struct {
	start, end, sum int
	left, right     *SegmentTree
}

func newSegmentTree(start, end int) *SegmentTree {
	return &SegmentTree{
		start: start,
		end:   end,
	}
}

func Constructor(nums []int) NumArray {
	return NumArray{root: buildTree(nums, 0, len(nums)-1)}
}

func buildTree(nums []int, start, end int) *SegmentTree {
	if start > end {
		return nil
	}

	ret := newSegmentTree(start, end)
	if start == end {
		ret.sum = nums[start]
	} else {
		mid := start + (end-start)/2
		ret.left = buildTree(nums, start, mid)
		ret.right = buildTree(nums, mid+1, end)
		ret.sum = ret.left.sum + ret.right.sum
	}

	return ret
}

func (this *NumArray) Update(index int, val int) {
	update(this.root, index, val)
}

func update(root *SegmentTree, pos, val int) {
	if root.start == root.end {
		root.sum = val
	} else {
		mid := root.start + (root.end-root.start)/2
		if pos <= mid {
			update(root.left, pos, val)
		} else {
			update(root.right, pos, val)
		}
		root.sum = root.left.sum + root.right.sum
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return sumRange(this.root, left, right)
}

func sumRange(root *SegmentTree, start, end int) int {
	if root.end == end && root.start == start {
		return root.sum
	}

	mid := root.start + (root.end-root.start)/2
	if end <= mid {
		return sumRange(root.left, start, end)
	} else if start >= mid+1 {
		return sumRange(root.right, start, end)
	}
	return sumRange(root.right, mid+1, end) + sumRange(root.left, start, mid)
}
