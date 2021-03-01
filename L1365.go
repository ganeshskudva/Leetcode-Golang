package Easy

func smallerNumbersThanCurrent(nums []int) []int {
	occurrences := [101]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	prevOccurrences := occurrences[0]
	occurrences[0] = 0
	for i := range occurrences[1:] {
		occurrences[i+1], prevOccurrences = prevOccurrences, prevOccurrences+occurrences[i+1]
	}
	counts := make([]int, len(nums))
	for i, num := range nums {
		counts[i] = occurrences[num]
	}
	return counts
}
