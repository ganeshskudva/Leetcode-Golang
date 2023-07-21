package Medium

type Pair struct {
	l, c int
}

func newPair(a, b int) *Pair {
	return &Pair{
		l: a,
		c: b,
	}
}

func findNumberOfLIS(nums []int) int {
	var (
		solve      func(idx int) *Pair
		maxL, maxC int
	)
	dp := make(map[int]*Pair)

	solve = func(idx int) *Pair {
		if idx == len(nums) {
			return newPair(0, 0)
		}
		if v, ok := dp[idx]; ok {
			return v
		}

		ml, mc := 0, 0
		for i := idx + 1; i < len(nums); i++ {
			if nums[i] > nums[idx] {
				tmp := solve(i)
				if tmp.l > ml {
					ml = tmp.l
					mc = tmp.c
				} else if tmp.l == ml {
					mc += tmp.c
				}
			}
		}

		if ml == 0 && mc == 0 {
			dp[idx] = newPair(1, 1)
			return dp[idx]
		}

		dp[idx] = newPair(1+ml, mc)
		return dp[idx]
	}

	for i := range nums {
		tmpPair := solve(i)
		if tmpPair.l > maxL {
			maxL = tmpPair.l
			maxC = tmpPair.c
		} else if tmpPair.l == maxL {
			maxC += tmpPair.c
		}
	}

	return maxC
}
