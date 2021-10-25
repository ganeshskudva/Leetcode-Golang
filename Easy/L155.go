package Easy

type StackNode struct {
	val  int
	min  int
	next *StackNode
}

func NewStackNode(val, min int, next *StackNode) *StackNode {
	return &StackNode{
		val:  val,
		min:  min,
		next: next,
	}
}

type MinStack struct {
	head *StackNode
}

func Constructor() MinStack {
	return MinStack{head: nil}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = NewStackNode(val, val, nil)
	} else {
		min := val
		if this.head.min < min {
			min = this.head.min
		}
		this.head = NewStackNode(val, min, this.head)
	}
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}
