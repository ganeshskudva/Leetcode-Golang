package Hard

package main

func main() {

	n := []int{4, -2, 2, -4}
	println(canReorderDoubled(n))
}

func rectangleArea(rectangles [][]int) int {
	mod, res := 1000000007, 0
	var recList [][]int
	for _, r := range rectangles {
		addRectangle(&recList, &r, 0)
	}

	for _, r := range recList {
		res = res + ((r[2]-r[0])*(r[3]-r[1]))%mod
	}

	return res % mod
}

func addRectangle(recList *[][]int, curRec *[]int, start int) {
	if start >= len(*recList) {
		*recList = append(*recList, *curRec)
		return
	}

	r := (*recList)[start]

	if (*curRec)[2] <= r[0] || (*curRec)[3] <= r[1] || (*curRec)[0] >= r[2] || (*curRec)[1] >= r[3] {
		addRectangle(recList, curRec, start+1)
		return
	}

	if (*curRec)[0] < r[0] {
		tmp := []int{(*curRec)[0], (*curRec)[1], r[0], (*curRec)[3]}
		addRectangle(recList, &tmp, start+1)
	}

	if (*curRec)[2] > r[2] {
		tmp := []int{r[2], (*curRec)[1], (*curRec)[2], (*curRec)[3]}
		addRectangle(recList, &tmp, start+1)
	}

	if (*curRec)[1] < r[1] {
		tmp := []int{max(r[0], (*curRec)[0]), (*curRec)[1], min(r[2], (*curRec)[2]), r[1]}
		addRectangle(recList, &tmp, start+1)
	}

	if (*curRec)[3] > r[3] {
		tmp := []int{max(r[0], (*curRec)[0]), r[3], min(r[2], (*curRec)[2]), (*curRec)[3]}
		addRectangle(recList, &tmp, start+1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
