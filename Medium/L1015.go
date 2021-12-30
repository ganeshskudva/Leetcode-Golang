package Medium

// Time  Complexity:O(k)
// Space Complexity:O(k)
var exists struct{}

func smallestRepunitDivByK(k int) int {
	n, ans, mp := 0, 0, make(map[int]struct{})

	for {
		ans, n = ans+1, (n*10+1)%k
		if n == 0 {
			return ans
		}
		if _, ok := mp[n]; ok {
			return -1
		}
		mp[n] = exists
	}
}

// Time  Complexity:O(k)
// Space Complexity:O(1)
func smallestRepunitDivByK(k int) int {
	n := 0
	for i := 0; i < k; i++ {
		n = (n*10 + 1) % k
		if n == 0 {
			return i + 1
		}
	}

	return -1
}
