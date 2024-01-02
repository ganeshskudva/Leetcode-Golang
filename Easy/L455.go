package Easy

func findContentChildren(children []int, cookies []int) int {
	sort.Ints(children)
	sort.Ints(cookies)

	child := 0
	for cookie := 0; child < len(children) && cookie < len(cookies); cookie++ {
		if cookies[cookie] >= children[child] {
			child++
		}
	}
	
	return child
}
