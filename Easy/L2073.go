package Easy

func timeRequiredToBuy(tickets []int, k int) int {
	total := 0
	for i := 0; i < len(tickets); i++ {
		if tickets[i] <= tickets[k] {
			total += tickets[i]
		} else {
			total += tickets[k]
		}
		
		if i > k && tickets[i] >= tickets[k] {
			total--
		}
	}
	
	return total
}
