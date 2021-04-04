package Medium

func removeStones(stones [][]int) int {
	n := len(stones)
	parent := make([]int, n)
	for i:= 0; i < n; i++ {
		parent[i] = i
	}
	
	for i:= 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			s1, s2 := stones[i], stones[j]
			if s1[0] == s2[0] || s1[1] == s2[1] {
				pi, pj := find(parent, i), find(parent, j)
				parent[pj] = pi
			}
		}
	}
	
	res := 0
	for i := 0 ; i < n; i++ {
		if parent[i] == i {
			res++
		}
	}
	
	return n - res
 }
 
 func find(parent []int, idx int) int {
 	i := idx
 	for parent[i] != i {
 		i = parent[i]
	}
	parent[idx] = i
	return i
 }
