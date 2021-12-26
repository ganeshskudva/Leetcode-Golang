package Medium

type priorityQueue [][]int

func NewPriorityQueue() *priorityQueue {
	return &priorityQueue{}
}

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	x, y := pq.distance(pq[i][0], pq[i][1]), pq.distance(pq[j][0], pq[j][1])

	return x > y
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *priorityQueue) Pop() interface{} {
	val := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]

	return val
}

func (pq priorityQueue) distance(x, y int) int {
	return x*x + y*y
}

func kClosest(points [][]int, k int) [][]int {
	var res [][]int
	if len(points) == 0 {
		return res
	}

	pq := NewPriorityQueue()
	heap.Init(pq)

	for _, p := range points {
		heap.Push(pq, p)
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	for pq.Len() > 0 {
		res = append(res, heap.Pop(pq).([]int))
	}

	return res
}
