package Medium

type BrowserHistory struct {
	History *[]string
	Future  *[]string
}

func Constructor(homepage string) BrowserHistory {
	bh := BrowserHistory{
		History: new([]string),
		Future:  new([]string),
	}
	push(homepage, bh.History)
	clear(bh.Future)

	return bh
}

func (this *BrowserHistory) Visit(url string) {
	push(url, this.History)
	clear(this.Future)
}

func (this *BrowserHistory) Back(steps int) string {
	for steps > 0 && length(this.History) > 1 {
		push(pop(this.History), this.Future)
		steps--
	}
	
	return top(this.History)
}

func (this *BrowserHistory) Forward(steps int) string {
	for steps > 0 && length(this.Future) > 0 {
		push(pop(this.Future), this.History)
		steps--
	}
	
	return top(this.History)
}

func push(url string, st *[]string) {
	*st = append(*st, url)
}

func pop(st *[]string) string {
	if len(*st) == 0 {
		return ""
	}

	val := top(st)
	*st = (*st)[:len(*st)-1]

	return val
}

func top(st *[]string) string {
	if len(*st) == 0 {
		return ""
	}

	return (*st)[len(*st)-1]
}

func clear(st *[]string) {
	*st = (*st)[:0]
}

func length(st *[]string) int {
	return len(*st)
}
