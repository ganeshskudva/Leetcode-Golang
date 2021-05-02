package Hard

type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }

func scheduleCourse(courses [][]int) int {
	day, maxHeap := 0, &MaxHeap{}
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	for _, course := range courses {
		heap.Push(maxHeap, course[0])
		day += course[0]

		if day > course[1] {
			day -= heap.Pop(maxHeap).(int)
		}

		if day > course[1] {
			return maxHeap.Len()
		}
	}

	return maxHeap.Len()
}
