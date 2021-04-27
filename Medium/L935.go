package Medium

const mod = 1000000007

func knightDialer(n int) int {
	// All the different places knight can jump from given location on keypad
	// From 0 can jump to {4,6} , from 1 can jump to {8,6} ....
	graph := [][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {3, 9, 0}, {}, {1, 7, 0}, {2, 6}, {1, 3}, {2, 4}}
	cnt := 0
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 10)
		for j:= range dp[i] {
			dp[i][j] = -1
		}
	}
	
	for i:= 0; i <= 9; i++ {
		cnt = (cnt + solve(n-1, i, graph, &dp)) % mod
	}
	
	return cnt
}

func solve(n, curr int, graph [][]int, dp *[][]int) int{
	if n== 0 {
		return 1
	}
	
	if (*dp)[n][curr] != -1 {
		return (*dp)[n][curr]
	}
	
	cnt := 0
	for _, nei := range graph[curr] {
		cnt = (cnt + solve(n-1, nei, graph, dp)) % mod
	}
	(*dp)[n][curr] = cnt
	
	return cnt
}
