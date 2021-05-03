package Easy

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func (h *MinHeap) Top() int {
	return (*h)[0]
}

type KthLargest struct {
	K  int
	PQ *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	kthLargest := KthLargest{
		K:  k,
		PQ: h,
	}

	for _, v := range nums {
		heap.Push(h, v)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.PQ, val)

	if this.PQ.Len() > this.K {
		heap.Pop(this.PQ)
	}

	return this.PQ.Top()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
