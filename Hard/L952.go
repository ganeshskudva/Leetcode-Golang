package Hard

type UnionFind struct {
	Parent []int
	Size   []int
	Max    int
}

func NewUnionFind(n int) *UnionFind {
	parent, size, max := make([]int, n), make([]int, n), 1
	for i := 0; i < n; i++ {
		parent[i], size[i] = i, 1
	}

	return &UnionFind{
		Parent: parent,
		Size:   size,
		Max:    max,
	}
}

func (u *UnionFind) find(x int) int {
	if x == u.Parent[x] {
		return x
	}
	u.Parent[x] = u.find(u.Parent[x])
	return u.Parent[x]
}

func (u *UnionFind) union(x, y int) {
	rootX, rootY := u.find(x), u.find(y)

	if rootY != rootX {
		u.Parent[rootX] = rootY
		u.Size[rootY] += u.Size[rootX]
		if u.Size[rootY] > u.Max {
			u.Max = u.Size[rootY]
		}
	}
}

func largestComponentSize(nums []int) int {
	mp, uf := make(map[int]int), NewUnionFind(len(nums))

	for i := range nums {
		a := nums[i]
		for j := 2; j*j <= a; j++ {
			if a%j == 0 {
				if _, ok := mp[j]; !ok {
					mp[j] = i
				} else {
					uf.union(i, mp[j])
				}

				if _, ok := mp[a/j]; !ok {
					mp[a/j] = i
				} else {
					uf.union(i, mp[a/j])
				}
			}
		}
		if _, ok := mp[a]; !ok {
			mp[a] = i
		} else {
			uf.union(i, mp[a])
		}
	}

	return uf.Max
}
