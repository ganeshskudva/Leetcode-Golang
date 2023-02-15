package Easy

func addToArrayForm(num []int, k int) []int {
	var res []int

	for i := len(num) - 1; i >= 0; i-- {
		res = append([]int{(num[i] + k) % 10}, res...)
		k = (num[i] + k) / 10
	}
	
	for k > 0 {
		res = append([]int{k % 10}, res...)
		k /= 10
	}
	
	return res
}
