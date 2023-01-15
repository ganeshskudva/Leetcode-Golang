package Hard

type UnionFind struct {
	parent []int
}

func newUnionFind(n int) *UnionFind {
	par := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = -1
	}
	return &UnionFind{
		parent: par,
	}
}

func (u *UnionFind) find(x int) int {
	if u.parent[x] >= 0 {
		u.parent[x] = u.find(u.parent[x])
		return u.parent[x]
	}

	return x
}

func (u *UnionFind) union(x, y int) bool {
	parentX, parentY := u.find(x), u.find(y)

	if parentY == parentX {
		return false
	}

	if u.parent[parentX] <= u.parent[parentY] {
		u.parent[parentX] += u.parent[parentY]
		u.parent[parentY] = parentX
	} else {
		u.parent[parentY] += u.parent[parentX]
		u.parent[parentX] = parentY
	}

	return true
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	adjList, mp, n := make(map[int][]int), make(map[int][]int), len(vals)

	for i := range vals {
		mp[vals[i]] = append(mp[vals[i]], i)
	}

	for _, e := range edges {
		u, v := e[0], e[1]

		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	var keys []int
	for k, _ := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	uf, res := newUnionFind(n), n

	for _, k := range keys {
		for _, curr := range mp[k] {
			for _, adj := range adjList[curr] {
				if vals[adj] <= vals[curr] {
					uf.union(curr, adj)
				}
			}
		}
		
		group := make(map[int]int) 
		for _, curr := range mp[k] {
			group[uf.find(curr)]++
		}
		
		for _, v := range group {
			if v >= 2 {
				res += v * (v - 1) / 2
			}
		}
	}
	
	return res
}
