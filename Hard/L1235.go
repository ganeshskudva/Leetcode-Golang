package Hard

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	memo := make(map[int]int)
	jobs := make([][]int, len(startTime))
	for i := 0; i < len(startTime); i++ {
		jobs[i] = []int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	return dfs(0, jobs, memo)
}

func dfs(cur int, jobs [][]int, mp map[int]int) int {
	if cur == len(jobs) {
		return 0
	}

	if v, ok := mp[cur]; ok {
		return v
	}

	next := findNext(cur, jobs)
	inclProf := jobs[cur][2]
	if next != -1 {
		inclProf += dfs(next, jobs, mp)
	}
	exclProf := dfs(cur+1, jobs, mp)

	max := inclProf
	if exclProf > inclProf {
		max = exclProf
	}

	mp[cur] = max
	return max
}

func findNext(cur int, jobs [][]int) int {
	lo, hi := cur+1, len(jobs)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if jobs[mid][0] >= jobs[cur][1] {
			if jobs[mid-1][0] >= jobs[cur][1] {
				hi = mid - 1
			} else {
				return mid
			}
		} else {
			lo = mid + 1
		}
	}

	return -1
}
