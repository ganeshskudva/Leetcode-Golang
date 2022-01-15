package Hard

var exists struct{}

func minJumps(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	mp := make(map[int][]int)
	for i := range arr {
		if _, ok := mp[arr[i]]; !ok {
			mp[arr[i]] = make([]int, 0)
		}
		mp[arr[i]] = append(mp[arr[i]], i)
	}
	visited := make(map[int]struct{})
	var q []int
	q = append(q, 0)
	visited[0] = exists
	cnt := 0

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			if n == len(arr)-1 {
				return cnt
			}
			if n > 0 {
				if _, ok := visited[n-1]; !ok {
					visited[n-1] = exists
					q = append(q, n-1)
				}
			}
			if n < len(arr)-1 {
				if _, ok := visited[n+1]; !ok {
					visited[n+1] = exists
					q = append(q, n+1)
				}
			}
			if _, ok := mp[arr[n]]; ok {
				for _, v := range mp[arr[n]] {
					if _, present := visited[v]; !present {
						visited[v] = exists
						q = append(q, v)
					}
				}
				delete(mp, arr[n])
			}
		}
		cnt++
	}

	return -1
}
