package Medium

type Solution struct {
	ans []int
	dp  []int
	tmp []int
}

func largestDivisibleSubset(nums []int) []int {
	s := &Solution{
		dp:  make([]int, 0),
		tmp: make([]int, 0),
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		s.dp = append(s.dp, -1)
	}
	s.solve(0, 1, nums)

	return s.ans
}

func (s *Solution) solve(i, prev int, nums []int) {
	if i >= len(nums) {
		if len(s.tmp) > len(s.ans) {
			s.ans = make([]int, len(s.tmp))
			copy(s.ans, s.tmp)
		}
		return
	}

	if len(s.tmp) > s.dp[i] && nums[i]%prev == 0 {
		s.dp[i] = len(s.tmp)
		s.tmp = append(s.tmp, nums[i])
		s.solve(i+1, nums[i], nums)
		s.tmp = s.tmp[:len(s.tmp)-1]
	}
	s.solve(i+1, prev, nums)
}
