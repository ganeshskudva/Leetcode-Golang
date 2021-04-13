package Medium

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	arr []int
	idx int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var a []int
	dfs(&a, nestedList)
	return &NestedIterator{arr: a, idx: 0}
}

func dfs(a *[]int, nl []*NestedInteger) {
	if nl == nil {
		return
	}

	for i:= 0; i < len(nl); i++ {
		if nl[i].IsInteger(){
			*a = append(*a, nl[i].GetInteger())
		} else {
			dfs(a, nl[i].GetList())
		}
	}
}

func (this *NestedIterator) Next() int {
	val := -1

	if this.HasNext() {
		val = this.arr[this.idx]
		this.idx++
	}

	return val
}

func (this *NestedIterator) HasNext() bool {
	return this.idx < len(this.arr)
}
