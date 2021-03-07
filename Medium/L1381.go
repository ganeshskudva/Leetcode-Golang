package Medium 

type CustomStack struct {
	st   []int
	size int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{st: make([]int, 0, maxSize), size: maxSize}
}

func (this *CustomStack) Push(x int) {
	if len(this.st) < this.size {
		this.st = append(this.st, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.st) == 0 {
		return -1
	}

	elem := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	return elem
}

func (this *CustomStack) Increment(k int, val int) {
	cnt := 0
	for i := 0; i < len(this.st) && cnt < k; i, cnt = i+1, cnt+1 {
		this.st[i] += val
	}
}
