package Easy

func getImportance(employees []*Employee, id int) int {
	if len(employees) == 0 {
		return 0
	}

	mp := make(map[int]int)
	dict := make(map[int]*Employee)
	for _, e := range employees {
		mp[e.Id] = e.Importance
		dict[e.Id] = e
	}

	sum := 0
	q := make([]*Employee, 1, 10)
	q[0] = dict[id]

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			e := q[0]

			sum += e.Importance
			for _, sub := range e.Subordinates {
				q = append(q, dict[sub])
			}
		}
	}

	return sum
}
