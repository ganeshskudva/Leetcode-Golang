package Easy

type OrderedStream struct {
	arr []string
	ptr int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		arr: make([]string, n+1),
		ptr: 0,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.arr[idKey-1] = value

	var res []string
	for ; this.ptr < len(this.arr) && this.arr[this.ptr] != ""; this.ptr++ {
		res = append(res, this.arr[this.ptr])
	}

	return res
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
