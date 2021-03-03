package Easy

func missingNumber(nums []int) int {
  sum := 0;
  for _, num := range nums {
    sum += num
  }
  length := len(nums)
  total := length*(length+1)/2;
  return total-sum
}
