type Task struct {
	processTime int
	index       int
}

type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].processTime < h[j].processTime || (h[i].processTime == h[j].processTime && h[i].index < h[j].index) }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func getOrder(tasks [][]int) []int {
	dic := make(map[int][]Task)
	for i := 0; i < len(tasks); i++ {
		dic[tasks[i][0]] = append(dic[tasks[i][0]], Task{tasks[i][1], i})
	}

	ans := []int{}
	keys := []int{}
	for k := range dic {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for len(keys) > 0 {
		k := keys[0]
		keys = keys[1:]
		pq := TaskHeap(dic[k])
		heap.Init(&pq)
		time := k

		for pq.Len() > 0 {
			task := heap.Pop(&pq).(Task)
			ans = append(ans, task.index)
			time += task.processTime
			for len(keys) > 0 && keys[0] <= time {
				for _, item := range dic[keys[0]] {
					heap.Push(&pq, item)
				}
				keys = keys[1:]
			}
		}
	}

	return ans
}
