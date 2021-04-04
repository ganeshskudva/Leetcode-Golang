package Medium

type MyCircularQueue struct {
	Q        []int
	Capacity int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Q:        make([]int, 0, k),
		Capacity: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.Q = append(this.Q, value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.Q = this.Q[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Q[0]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Q[len(this.Q)-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.Q) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.Q) == this.Capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
