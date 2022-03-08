package Easy

func hasCycle(head *ListNode) bool {
    	walker, runner := head, head
	
	for runner != nil && runner.Next != nil {
		walker, runner = walker.Next, runner.Next.Next
		if walker == runner {
			return true
		}
	}
	
	return false
}
