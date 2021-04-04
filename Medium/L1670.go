package Medium

type FrontMiddleBackQueue struct {
	Q []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		Q: []int{},
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.Q = append([]int{val}, this.Q...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	middle := len(this.Q) / 2
	if len(this.Q) == middle {
		this.Q = append(this.Q, val)
        return
	}

	this.Q = append(this.Q[:middle+1], this.Q[middle:]...)
	this.Q[middle] = val
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.Q = append(this.Q, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.isEmpty() {
		return -1
	}

	ret := this.Q[0]
	this.Q = this.Q[1:]
	return ret
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.isEmpty() {
		return -1
	}

	middle := math.Ceil(float64(len(this.Q)) / 2) - 1
	ret := this.Q[int(middle)]
	this.Q = append(this.Q[:int(middle)], this.Q[int(middle)+1:]...)
	return ret
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.isEmpty() {
		return -1
	}

	ret := this.Q[len(this.Q)-1]
	this.Q = this.Q[:len(this.Q)-1]
	return ret
}

func (this *FrontMiddleBackQueue) isEmpty() bool {
	return len(this.Q) == 0
}


/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
