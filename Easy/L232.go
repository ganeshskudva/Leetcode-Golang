type MyQueue struct {
	q []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{q: make([]int, 0)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.q = append(this.q, x)
}

/** Removes the element from in front of queue and returns that element. */
package Easy

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	ret := this.q[0]
	this.q = this.q[1:]
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.q[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.q) == 0 {
		return true
	}

	return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
