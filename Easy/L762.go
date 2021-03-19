package Easy

func countPrimeSetBits(L int, R int) int {
    	mp := make(map[int]bool)
	val := []int{2,3,5,7,11,13,17,19}
	
	for _, v := range val {
		mp[v] = true
	}
	
	cnt := 0
	for i:= L; i <= R; i++ {
		bits := 0
		for n := i; n > 0; n >>= 1 {
			bits += n &1
		}
		if _, ok := mp[bits]; ok {
			cnt++
		} 
	}
	
	return cnt
}
